package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		LogLevel string `mapstructure:"log_level"`
	}

	Bot struct {
		Token string `mapstructure:"BOT_TOKEN"`
	} `mapstructure:",squash"`

	Mongo struct {
		URI string `mapstructure:"MONGO_URI"`
		DbName string `mapstructure:"MONGO_INITDB_DATABASE"`
		User string `mapstructure:"MONGO_INITDB_ROOT_USERNAME"`
		Password string `mapstructure:"MONGO_INITDB_ROOT_PASSWORD"`
	} `mapstructure:",squash"`
}

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
