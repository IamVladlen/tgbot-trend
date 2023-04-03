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

	Bot struct {
		Token string `mapstructure:"BOT_TOKEN"`
	} `mapstructure:",squash"`

	GRPC struct {
		URI string `mapstructure:"GRPC_SCHEDULER_URI"`
	} `mapstructure:",squash"`

	Mongo struct {
		URI      string `mapstructure:"MONGO_URI"`
		DBName   string `mapstructure:"MONGO_INITDB_DATABASE"`
		User     string `mapstructure:"MONGO_INITDB_ROOT_USERNAME"`
		Password string `mapstructure:"MONGO_INITDB_ROOT_PASSWORD"`
	} `mapstructure:",squash"`

	Redis struct {
		URI      string `mapstructure:"REDIS_URI"`
		Password string `mapstructure:"REDIS_PASSWORD"`
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
	cfg.Bot.Token = os.Getenv("BOT_TOKEN")

	cfg.GRPC.URI = os.Getenv("GRPC_SCHEDULER_URI")

	cfg.Mongo.URI = os.Getenv("MONGO_URI")
	cfg.Mongo.DBName = os.Getenv("MONGO_INITDB_DATABASE")
	cfg.Mongo.User = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	cfg.Mongo.Password = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")

	cfg.Redis.URI = os.Getenv("REDIS_URI")
	cfg.Redis.URI = os.Getenv("REDIS_PASSWORD")
}
