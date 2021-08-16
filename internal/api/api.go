package api

import (
	"context"

	"github.com/rs/zerolog/log"

	desc "github.com/ozoncp/ocp-survey-api/pkg/ocp-survey-api"
)

type api struct {
	desc.UnimplementedOcpSurveyApiServer
}

func NewOcpSurveyApi() desc.OcpSurveyApiServer {
	return &api{}
}

func (a *api) CreateSurveyV1(ctx context.Context, in *desc.CreateSurveyV1Request) (*desc.CreateSurveyV1Response, error) {
	log.Info().
		Uint64("user_id", in.GetUserId()).
		Str("link", in.GetLink()).
		Msg("Create survey request")
	return &desc.CreateSurveyV1Response{}, nil
}

func (a *api) DescribeSurveyV1(ctx context.Context, in *desc.DescribeSurveyV1Request) (*desc.DescribeSurveyV1Response, error) {
	log.Info().
		Uint64("survey_id", in.GetSurveyId()).
		Msg("Describe survey request")
	return &desc.DescribeSurveyV1Response{}, nil
}

func (a *api) ListSurveysV1(ctx context.Context, in *desc.ListSurveysV1Request) (*desc.ListSurveysV1Response, error) {
	log.Info().
		Msg("List surveys request")
	return &desc.ListSurveysV1Response{}, nil
}

func (a *api) RemoveSurveyV1(ctx context.Context, in *desc.RemoveSurveyV1Request) (*desc.RemoveSurveyV1Response, error) {
	log.Info().
		Uint64("survey_id", in.GetSurveyId()).
		Msg("Remove survey request")
	return &desc.RemoveSurveyV1Response{}, nil
}
