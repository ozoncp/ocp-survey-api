package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

// regSignalHandler отменяет контекст при получении сигналов SIGQUIT, SIGINT, SIGTERM.
func regSignalHandler(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer signal.Stop(stop)
		<-stop
		log.Info().Msg("Stop signal received")
		cancel()
	}()

	return ctx
}

// run запускает gRPC-сервер и JSON-gateway к нему.
// При отмене контекста ctx будет выполнен graceful stop.
func run(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error { return runGRPC(ctx) })
	g.Go(func() error { return runJSON(ctx) })

	return g.Wait()
}

func runGRPC(ctx context.Context) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Error().Err(err).Msg("GRPC: Listen")
		return err
	}

	srv := grpc.NewServer()
	desc.RegisterOcpSurveyApiServer(srv, api.NewOcpSurveyApi())

	srvErr := make(chan error)
	go func() {
		if err := srv.Serve(listen); err != nil {
			srvErr <- err
		}
	}()

	select {
	case err := <-srvErr:
		log.Error().Err(err).Msg("GRPC: Serve")
		return err

	case <-ctx.Done():
		srv.GracefulStop()
	}

	return nil
}

func runJSON(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterOcpSurveyApiHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		log.Error().Err(err).Msg("JSON: Register API handler")
		return err
	}

	srv := &http.Server{Addr: httpPort, Handler: mux}

	srvErr := make(chan error)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			srvErr <- err
		}
	}()

	select {
	case err := <-srvErr:
		log.Error().Err(err).Msg("JSON: Serve")
		return err

	case <-ctx.Done():
		err := srv.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("JSON: Shutdown")
			return err
		}
	}

	return nil
}

func main() {
	log.Info().Msg("Ozon Code Platform Survey service started")

	ctx := regSignalHandler(context.Background())
	if err := run(ctx); err != nil {
		log.Fatal().Err(err).Msg("Service stopped on error")
	}
	log.Info().Msg("Service exited")
}
