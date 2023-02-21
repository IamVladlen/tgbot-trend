package config

import (
	"log"

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

	viper.SetConfigFile("./.env")
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalln("Can't load config file:", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalln("Can't load config file:", err)
	}

	return cfg
}
