package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// TODO: Set enviroment variables via github.com/kelseyhightower/envconfig

// Config stores data from ./config/config.yaml and ./.env files.
type Config struct {
	App struct {
		LogLevel string `mapstructure:"log_level"`
	}

	GRPC struct {
		Port string `mapstructure:"GRPC_SCHEDULER_PORT"`
	} `mapstructure:",squash"`

	PG struct {
		URL       string `mapstructure:"POSTGRES_URL"`
		Migration string `mapstructure:"PG_MIGRATION_URL"`
	} `mapstructure:",squash"`
}

// New creates new config instance with data from
// config files.
func New() *Config {
	var cfg *Config

	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Can't load config file:", err)
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalln("Can't load config file:", err)
	}

	loadEnv(cfg)

	return cfg
}

func loadEnv(cfg *Config) {
	cfg.GRPC.Port = os.Getenv("GRPC_SCHEDULER_PORT")

	cfg.PG.URL = os.Getenv("POSTGRES_URL")
	cfg.PG.Migration = os.Getenv("PG_MIGRATION_URL")
}
