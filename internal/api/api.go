package api

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozoncp/ocp-survey-api/internal/models"
	"github.com/ozoncp/ocp-survey-api/internal/repo"
	"github.com/ozoncp/ocp-survey-api/internal/utils"
	desc "github.com/ozoncp/ocp-survey-api/pkg/ocp-survey-api"
)

type api struct {
	desc.UnimplementedOcpSurveyApiServer
	repo      repo.Repo
	chunkSize int
}

func NewOcpSurveyApi(repo repo.Repo, chunkSize int) desc.OcpSurveyApiServer {
	return &api{
		repo:      repo,
		chunkSize: chunkSize,
	}
}

func (a *api) CreateSurveyV1(ctx context.Context, in *desc.CreateSurveyV1Request) (*desc.CreateSurveyV1Response, error) {
	log.Info().
		Uint64("user_id", in.GetUserId()).
		Str("link", in.GetLink()).
		Msg("Create survey request")

	survey := models.Survey{
		UserId: in.GetUserId(),
		Link:   in.GetLink(),
	}

	ids, err := a.repo.AddSurvey(ctx, []models.Survey{survey})
	if err != nil {
		log.Error().Err(err).Msg("Create survey: failed")
		return nil, status.Error(codes.Internal, "internal error")
	}

	res := &desc.CreateSurveyV1Response{}
	if len(ids) > 0 {
		res.SurveyId = ids[0]
	}

	return res, nil
}

func (a *api) MultiCreateSurveyV1(ctx context.Context, in *desc.MultiCreateSurveyV1Request) (*desc.MultiCreateSurveyV1Response, error) {
	log.Info().
		Int("num_items", len(in.GetSurveys())).
		Msg("Multi create survey request")

	inSurveys := in.GetSurveys()
	if len(inSurveys) == 0 {
		log.Error().Msg("Multi create survey: no surveys passed")
		return nil, status.Error(codes.InvalidArgument, "no surveys to store")
	}

	surveys := make([]models.Survey, 0, len(inSurveys))
	for _, item := range inSurveys {
		survey := models.Survey{
			UserId: item.GetUserId(),
			Link:   item.GetLink(),
		}
		surveys = append(surveys, survey)
	}

	ids := make([]uint64, 0, len(surveys))
	chunks, err := utils.SplitToChunks(surveys, a.chunkSize)
	if err != nil {
		log.Error().Err(err).Msg("Multi create survey: split to chunks failed")
		return nil, status.Error(codes.Internal, "internal error")
	}

	for idx, chunk := range chunks {
		newIds, err := a.repo.AddSurvey(ctx, chunk)
		if err != nil {
			log.Error().
				Int("chunk", idx).
				Err(err).
				Msg("Multi create survey: failed")
			res := &desc.MultiCreateSurveyV1Response{
				SurveyIds: ids,
			}
			return res, status.Error(codes.Internal, "internal error")
		}
		ids = append(ids, newIds...)
	}

	res := &desc.MultiCreateSurveyV1Response{
		SurveyIds: ids,
	}
	return res, nil
}

func (a *api) DescribeSurveyV1(ctx context.Context, in *desc.DescribeSurveyV1Request) (*desc.DescribeSurveyV1Response, error) {
	log.Info().
		Uint64("survey_id", in.GetSurveyId()).
		Msg("Describe survey request")

	survey, err := a.repo.DescribeSurvey(ctx, in.GetSurveyId())
	if errors.Is(err, repo.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "survey id not found")
	} else if err != nil {
		log.Error().Err(err).Msg("Describe survey: failed")
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
	log.Info().
		Uint64("limit", in.GetLimit()).
		Uint64("offset", in.GetOffset()).
		Msg("List surveys request")

	surveys, err := a.repo.ListSurveys(ctx, in.GetLimit(), in.GetOffset())
	if err != nil {
		log.Error().Err(err).Msg("List surveys: failed")
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
	inSurvey := in.GetSurvey()

	log.Info().
		Uint64("survey_id", inSurvey.GetId()).
		Uint64("user_id", inSurvey.GetUserId()).
		Str("link", inSurvey.GetLink()).
		Msg("Update survey request")

	survey := models.Survey{
		Id:     inSurvey.GetId(),
		UserId: inSurvey.GetUserId(),
		Link:   inSurvey.GetLink(),
	}

	err := a.repo.UpdateSurvey(ctx, survey)
	if errors.Is(err, repo.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "survey id not found")
	} else if err != nil {
		log.Error().Err(err).Msg("Describe survey: failed")
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &desc.UpdateSurveyV1Response{}, nil
}

func (a *api) RemoveSurveyV1(ctx context.Context, in *desc.RemoveSurveyV1Request) (*desc.RemoveSurveyV1Response, error) {
	log.Info().
		Uint64("survey_id", in.GetSurveyId()).
		Msg("Remove survey request")

	err := a.repo.RemoveSurvey(ctx, in.GetSurveyId())
	if err != nil {
		if err == repo.ErrNotFound {
			return nil, status.Error(codes.NotFound, "survey id not found")
		}
		log.Error().Err(err).Msg("Remove survey: failed")
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &desc.RemoveSurveyV1Response{}, nil
}
