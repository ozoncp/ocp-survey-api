package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog/log"
)

const configFileName = "config.yml"

type Config struct {
	ChunkSize int    `yaml:"chunksize" env:"SURVEY_CHUNK_SIZE" env-default:"32"`
	LogLevel  string `yaml:"loglevel" env:"SURVEY_LOG_LEVEL"`

	GRPC struct {
		Bind string `yaml:"bind" env:"SURVEY_GRPC_BIND"`
		Port string `yaml:"port" env:"SURVEY_GRPC_PORT" env-default:"8082"`
	} `yaml:"grpc"`

	Gateway struct {
		Bind string `yaml:"bind" env:"SURVEY_GW_BIND"`
		Port string `yaml:"port" env:"SURVEY_GW_PORT" env-default:"8080"`
	} `yaml:"gateway"`

	Database struct {
		Host string `yaml:"host" env:"SURVEY_DB_HOST" env-default:"localhost"`
		Port string `yaml:"port" env:"SURVEY_DB_PORT" env-default:"5432"`
		User string `yaml:"user" env:"SURVEY_DB_USER" env-default:"postgres"`
		Pass string `yaml:"pass" env:"SURVEY_DB_PASS" env-default:"postgres"`
		Name string `yaml:"name" env:"SURVEY_DB_NAME" env-default:"postgres"`
	} `yaml:"database"`

	Metrics struct {
		Bind string `yaml:"bind" env:"SURVEY_METRICS_BIND"`
		Port string `yaml:"port" env:"SURVEY_METRICS_PORT" env-default:"9100"`
	} `yaml:"metrics"`

	Broker struct {
		List []string `yaml:"list" env:"SURVEY_BROKERS" env-default:"localhost:9094"`
	} `yaml:"broker"`

	Tracing struct {
		AgentHost string `yaml:"agent_host" env:"SURVEY_TRACING_AGENT_HOST" env-default:"localhost"`
		AgentPort string `yaml:"agent_port" env:"SURVEY_TRACING_AGENT_PORT" env-default:"6831"`
	} `yaml:"tracing"`
}

var cfg *Config

func init() {
	cfg = &Config{}
	err := cleanenv.ReadConfig(configFileName, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Load configuration")
	}
	log.Info().Msg("Configuration loaded")
}

func Get() *Config {
	return cfg
}
