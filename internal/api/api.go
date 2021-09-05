package api

import (
	"context"
	"errors"

	"github.com/opentracing/opentracing-go"
	tracerlog "github.com/opentracing/opentracing-go/log"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozoncp/ocp-survey-api/internal/metrics"
	"github.com/ozoncp/ocp-survey-api/internal/models"
	"github.com/ozoncp/ocp-survey-api/internal/producer"
	"github.com/ozoncp/ocp-survey-api/internal/repo"
	"github.com/ozoncp/ocp-survey-api/internal/utils"
	desc "github.com/ozoncp/ocp-survey-api/pkg/ocp-survey-api"
)

type api struct {
	desc.UnimplementedOcpSurveyApiServer
	repo   repo.Repo
	prod   producer.Producer
	metr   metrics.Metrics
	tracer opentracing.Tracer
	chunk  int
}

func NewOcpSurveyApi(
	repo repo.Repo,
	prod producer.Producer,
	metr metrics.Metrics,
	tracer opentracing.Tracer,
	chunkSize int,
) desc.OcpSurveyApiServer {
	return &api{
		repo:   repo,
		prod:   prod,
		metr:   metr,
		tracer: tracer,
		chunk:  chunkSize,
	}
}

const (
	logMethodKey  = "method_name"
	producerTopic = "survey_events"
)

func (a *api) reportEvent(typ producer.EventType, survey_id uint64) {
	ev := producer.PrepareEvent(typ, survey_id)
	err := a.prod.Send(producerTopic, ev)
	if err != nil {
		log.Error().Err(err).Msg("Producer: Send event")
	}
}

func (a *api) CreateSurveyV1(ctx context.Context, in *desc.CreateSurveyV1Request) (*desc.CreateSurveyV1Response, error) {
	methodName := "CreateSurveyV1"
	log.Info().
		Str(logMethodKey, methodName).
		Uint64("user_id", in.GetUserId()).
		Str("link", in.GetLink()).
		Msg("Create survey request")

	span := a.tracer.StartSpan(methodName)
	defer span.Finish()

	if err := in.Validate(); err != nil {
		log.Error().Err(err).Str(logMethodKey, methodName).Msg("validation failed")
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	survey := models.Survey{
		UserId: in.GetUserId(),
		Link:   in.GetLink(),
	}

	ids, err := a.repo.AddSurvey(ctx, []models.Survey{survey})
	if err != nil {
		log.Error().Err(err).Str(logMethodKey, methodName).Msg("repo: add failed")
		return nil, status.Error(codes.Internal, "internal error")
	}

	res := &desc.CreateSurveyV1Response{}
	if len(ids) > 0 {
		res.SurveyId = ids[0]
		a.reportEvent(producer.Create, ids[0])
		a.metr.IncCreate()
	}

	return res, nil
}

func (a *api) MultiCreateSurveyV1(ctx context.Context, in *desc.MultiCreateSurveyV1Request) (*desc.MultiCreateSurveyV1Response, error) {
	methodName := "MultiCreateSurveyV1"
	log.Info().
		Str(logMethodKey, methodName).
		Int("num_items", len(in.GetSurveys())).
		Msg("Multi create survey request")

	span := a.tracer.StartSpan(methodName)
	defer span.Finish()

	if err := in.Validate(); err != nil {
		log.Error().Err(err).Str(logMethodKey, methodName).Msg("validation failed")
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	inSurveys := in.GetSurveys()
	surveys := make([]models.Survey, 0, len(inSurveys))
	for _, item := range inSurveys {
		survey := models.Survey{
			UserId: item.GetUserId(),
			Link:   item.GetLink(),
		}
		surveys = append(surveys, survey)
	}

	ids := make([]uint64, 0, len(surveys))
	chunks, err := utils.SplitToChunks(surveys, a.chunk)
	if err != nil {
		log.Error().Err(err).Str(logMethodKey, methodName).Msg("split to chunks failed")
		return nil, status.Error(codes.Internal, "internal error")
	}

	addChunk := func(ctx context.Context, surveys []models.Survey) ([]uint64, error) {
		chunkSpan := a.tracer.StartSpan("Chunk", opentracing.ChildOf(span.Context()))
		chunkSpan.LogFields(tracerlog.Int("num_items", len(surveys)))
		defer chunkSpan.Finish()
		return a.repo.AddSurvey(ctx, surveys)
	}

	for idx, chunk := range chunks {
		newIds, err := addChunk(ctx, chunk)
		if err != nil {
			log.Error().
				Str(logMethodKey, methodName).
				Int("chunk", idx).
				Err(err).
				Msg("repo: add chunk failed")
			res := &desc.MultiCreateSurveyV1Response{
				SurveyIds: ids,
			}
			return res, status.Error(codes.Internal, "internal error")
		}
		ids = append(ids, newIds...)

		for _, id := range newIds {
			a.reportEvent(producer.Create, id)
			a.metr.IncCreate()
		}
	}

	res := &desc.MultiCreateSurveyV1Response{
		SurveyIds: ids,
	}
	return res, nil
}

func (a *api) DescribeSurveyV1(ctx context.Context, in *desc.DescribeSurveyV1Request) (*desc.DescribeSurveyV1Response, error) {
	methodName := "DescribeSurveyV1"
	log.Info().
		Str(logMethodKey, methodName).
		Uint64("survey_id", in.GetSurveyId()).
		Msg("Describe survey request")

	span := a.tracer.StartSpan(methodName)
	defer span.Finish()

	if err := in.Validate(); err != nil {
		log.Error().Err(err).Str(logMethodKey, methodName).Msg("validation failed")
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	survey, err := a.repo.DescribeSurvey(ctx, in.GetSurveyId())
	if errors.Is(err, repo.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "survey id not found")
	} else if err != nil {
		log.Error().Err(err).Str(logMethodKey, methodName).Msg("repo: describe failed")
		return nil, status.Error(codes.Internal, "internal error")
	}

	res := &desc.DescribeSurveyV1Response{Survey: &desc.Survey{
		Id:     survey.Id,
		UserId: survey.UserId,
		Link:   survey.Link,
	}}

	return res, nil
}

func (a *api) ListSurveysV1(ctx context.Context, in *desc.ListSurveysV1Request) (*desc.ListSurveysV1Response, error) {
	methodName := "ListSurveysV1"
	log.Info().
		Str(logMethodKey, methodName).
		Uint64("limit", in.GetLimit()).
		Uint64("offset", in.GetOffset()).
		Msg("List surveys request")

	span := a.tracer.StartSpan(methodName)
	defer span.Finish()

	if err := in.Validate(); err != nil {
		log.Error().Err(err).Str(logMethodKey, methodName).Msg("validation failed")
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	surveys, err := a.repo.ListSurveys(ctx, in.GetLimit(), in.GetOffset())
	if err != nil {
		log.Error().Err(err).Str(logMethodKey, methodName).Msg("repo: list failed")
		return nil, status.Error(codes.Internal, "internal error")
	}

	res := &desc.ListSurveysV1Response{
		Surveys: make([]*desc.Survey, len(surveys)),
	}
	for idx, survey := range surveys {
		res.Surveys[idx] = &desc.Survey{
			Id:     survey.Id,
			UserId: survey.UserId,
			Link:   survey.Link,
		}
	}

	return res, nil
}

func (a *api) UpdateSurveyV1(ctx context.Context, in *desc.UpdateSurveyV1Request) (*desc.UpdateSurveyV1Response, error) {
	methodName := "UpdateSurveyV1"
	inSurvey := in.GetSurvey()

	log.Info().
		Str(logMethodKey, methodName).
		Uint64("survey_id", inSurvey.GetId()).
		Uint64("user_id", inSurvey.GetUserId()).
		Str("link", inSurvey.GetLink()).
		Msg("Update survey request")

	span := a.tracer.StartSpan(methodName)
	defer span.Finish()

	if err := in.Validate(); err != nil {
		log.Error().Err(err).Str(logMethodKey, methodName).Msg("validation failed")
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	survey := models.Survey{
		Id:     inSurvey.GetId(),
		UserId: inSurvey.GetUserId(),
		Link:   inSurvey.GetLink(),
	}

	err := a.repo.UpdateSurvey(ctx, survey)
	if errors.Is(err, repo.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "survey id not found")
	} else if err != nil {
		log.Error().Err(err).Str(logMethodKey, methodName).Msg("repo: update failed")
		return nil, status.Error(codes.Internal, "internal error")
	}

	a.reportEvent(producer.Update, survey.Id)
	a.metr.IncUpdate()

	return &desc.UpdateSurveyV1Response{}, nil
}

func (a *api) RemoveSurveyV1(ctx context.Context, in *desc.RemoveSurveyV1Request) (*desc.RemoveSurveyV1Response, error) {
	methodName := "RemoveSurveyV1"
	log.Info().
		Str(logMethodKey, methodName).
		Uint64("survey_id", in.GetSurveyId()).
		Msg("Remove survey request")

	span := a.tracer.StartSpan(methodName)
	defer span.Finish()

	if err := in.Validate(); err != nil {
		log.Error().Err(err).Str(logMethodKey, methodName).Msg("validation failed")
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	err := a.repo.RemoveSurvey(ctx, in.GetSurveyId())
	if errors.Is(err, repo.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "survey id not found")
	} else if err != nil {
		log.Error().Err(err).Str(logMethodKey, methodName).Msg("repo: remove failed")
		return nil, status.Error(codes.Internal, "internal error")
	}

	a.reportEvent(producer.Delete, in.GetSurveyId())
	a.metr.IncDelete()

	return &desc.RemoveSurveyV1Response{}, nil
}
