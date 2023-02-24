package main

import (
	"github.com/IamVladlen/trend-bot/config"
	"github.com/IamVladlen/trend-bot/internal/bot"
	"github.com/IamVladlen/trend-bot/pkg/logger"
)

func main() {
	cfg := config.New()
	log := logger.New(cfg.App.LogLevel)

	bot.Run(cfg, log)
}
