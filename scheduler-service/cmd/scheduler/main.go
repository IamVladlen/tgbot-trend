package main

import (
	"github.com/IamVladlen/trend-bot/scheduler-service/config"
	"github.com/IamVladlen/trend-bot/scheduler-service/pkg/logger"
	"github.com/IamVladlen/trend-bot/scheduler-service/internal/app"
)

func main() {
	cfg := config.New()
	log := logger.New(cfg.App.LogLevel)

	app.Run(log, cfg)
}