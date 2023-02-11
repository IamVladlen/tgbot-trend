package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

func New(logLevel string) *Logger {
	logger := &Logger{
		zerolog.New(os.Stdout).With().Timestamp().Logger().Level(setLogLevel(logLevel)),
	}

	return logger
}

func setLogLevel(logLevel string) zerolog.Level {
	switch logLevel {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	default:
		fmt.Printf("Incorrect log level: edit config.yaml\n")
		return zerolog.InfoLevel
	}
}
