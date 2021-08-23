package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-survey-api/internal/api"
	"github.com/ozoncp/ocp-survey-api/internal/repo"
	desc "github.com/ozoncp/ocp-survey-api/pkg/ocp-survey-api"
)

const (
	grpcPort     = ":8082"
	httpPort     = ":8080"
	grpcEndpoint = "localhost" + grpcPort

	dbUser = "postgres"
	dbPass = "postgres"
	dbHost = "localhost"
	dbPort = "5432"
	dbName = "postgres"

	chunkSize = 32
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

// run запускает сервис.
// При отмене контекста ctx будет выполнен graceful stop.
func run(ctx context.Context) error {
	repo, err := getRepo()
	if err != nil {
		return err
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error { return runGRPC(ctx, repo) })
	g.Go(func() error { return runJSON(ctx) })

	return g.Wait()
}

func getRepo() (repo.Repo, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Error().Err(err).Msg("DB: Connect")
		return nil, err
	}

	repo := repo.NewSurveyRepo(db)
	return repo, nil
}

func runGRPC(ctx context.Context, repo repo.Repo) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Error().Err(err).Msg("GRPC: Listen")
		return err
	}

	srv := grpc.NewServer()
	desc.RegisterOcpSurveyApiServer(srv, api.NewOcpSurveyApi(repo, chunkSize))

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
