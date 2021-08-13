package main

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-survey-api/internal/api"
	desc "github.com/ozoncp/ocp-survey-api/pkg/ocp-survey-api"
)

const (
	grpcPort     = ":8082"
	httpPort     = ":8080"
	grpcEndpoint = "localhost" + grpcPort
)

func run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error { return runGRPC(ctx) })
	g.Go(func() error { return runJSON(ctx) })

	return g.Wait()
}

func runGRPC(_ context.Context) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal().Msgf("GRPC Listen: %v", err)
		return err
	}

	s := grpc.NewServer()
	desc.RegisterOcpSurveyApiServer(s, api.NewOcpSurveyApi())

	err = s.Serve(listen)
	if err != nil {
		log.Fatal().Msgf("GRPC Serve: %v", err)
	}
	return err
}

func runJSON(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterOcpSurveyApiHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		log.Fatal().Msgf("Register JSON API handler: %v", err)
		return err
	}

	err = http.ListenAndServe(httpPort, mux)
	if err != nil {
		log.Fatal().Msgf("JSON Serve: %v", err)
	}
	return err
}

func main() {
	log.Info().Msg("Ozon Code Platform Survey service started")

	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Fatal().Msgf("Run service: %v", err)
	}
}
