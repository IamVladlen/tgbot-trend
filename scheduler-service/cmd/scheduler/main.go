package main

import (
	"github.com/IamVladlen/trend-bot/scheduler-service/config"
	"github.com/IamVladlen/trend-bot/scheduler-service/internal/app"
	"github.com/IamVladlen/trend-bot/scheduler-service/pkg/logger"
)

func main() {
	cfg := config.New()
	log := logger.New(cfg.App.LogLevel)

	app.Run(log, cfg)
}
