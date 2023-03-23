package main

import (
	"github.com/IamVladlen/trend-bot/bot-gateway/config"
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/bot"
	"github.com/IamVladlen/trend-bot/bot-gateway/pkg/logger"
)

func main() {
	cfg := config.New()
	log := logger.New(cfg.App.LogLevel)

	bot.Run(cfg, log)
}
