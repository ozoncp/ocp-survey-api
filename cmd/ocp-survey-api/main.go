package main

import (
	"context"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegermetrics "github.com/uber/jaeger-lib/metrics"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/ozoncp/ocp-survey-api/internal/api"
	"github.com/ozoncp/ocp-survey-api/internal/metrics"
	"github.com/ozoncp/ocp-survey-api/internal/producer"
	"github.com/ozoncp/ocp-survey-api/internal/repo"
	desc "github.com/ozoncp/ocp-survey-api/pkg/ocp-survey-api"
)

var (
	grpcPort        = ":8082"
	httpPort        = ":8080"
	dsn             = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	chunkSize       = 32
	brokers         = []string{"localhost:9094"}
	metricsPort     = ":9100"
	tracingHostPort = "localhost:6831"
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

	prod, err := producer.New(brokers)
	if err != nil {
		log.Error().Err(err).Msg("Producer: New")
		return err
	}
	defer prod.Close()

	metr := metrics.New()

	tracer, tracingCloser, err := initTracing()
	if err != nil {
		log.Error().Err(err).Msg("Tracing: Init")
		return err
	}
	defer tracingCloser.Close()

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error { return runService(ctx, repo, prod, metr, tracer) })
	g.Go(func() error { return runGateway(ctx) })
	g.Go(func() error { return runMetrics(ctx) })

	return g.Wait()
}

func getRepo() (repo.Repo, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Error().Err(err).Msg("DB: Connect")
		return nil, err
	}

	repo := repo.NewSurveyRepo(db)
	return repo, nil
}

func runService(
	ctx context.Context,
	repo repo.Repo,
	prod producer.Producer,
	metr metrics.Metrics,
	tracer opentracing.Tracer,
) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Error().Err(err).Msg("GRPC: Listen")
		return err
	}

	srv := grpc.NewServer()
	desc.RegisterOcpSurveyApiServer(srv, api.NewOcpSurveyApi(repo, prod, metr, tracer, chunkSize))

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

func runGateway(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	grpcEndpoint := "localhost" + grpcPort
	err := desc.RegisterOcpSurveyApiHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		log.Error().Err(err).Msg("Gateway: Register API handler")
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
		log.Error().Err(err).Msg("Gateway: Serve")
		return err

	case <-ctx.Done():
		err := srv.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("Gateway: Shutdown")
			return err
		}
	}

	return nil
}

func runMetrics(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{Addr: metricsPort, Handler: mux}

	srvErr := make(chan error)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			srvErr <- err
		}
	}()

	select {
	case err := <-srvErr:
		log.Error().Err(err).Msg("Metrics: Serve")
		return err

	case <-ctx.Done():
		err := srv.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("Metrics: Shutdown")
			return err
		}
	}

	return nil
}

func initTracing() (opentracing.Tracer, io.Closer, error) {
	cfg := jaegercfg.Configuration{
		ServiceName: "ocp-survey-api",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LocalAgentHostPort: tracingHostPort,
			LogSpans:           true,
		},
	}

	logger := jaeger.StdLogger
	metricsFactory := jaegermetrics.NullFactory

	return cfg.NewTracer(
		jaegercfg.Logger(logger),
		jaegercfg.Metrics(metricsFactory),
	)
}

func main() {
	log.Info().Msg("Ozon Code Platform Survey service started")

	ctx := regSignalHandler(context.Background())
	if err := run(ctx); err != nil {
		log.Fatal().Err(err).Msg("Service stopped on error")
	}
	log.Info().Msg("Service exited")
}
