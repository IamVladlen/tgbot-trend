package handler

import (
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/usecase"
	"github.com/IamVladlen/trend-bot/bot-gateway/pkg/logger"
	th "github.com/mymmrac/telego/telegohandler"
)

func New(handler *th.BotHandler, uc *usecase.UseCase, log *logger.Logger) *th.BotHandler {
	// Initialize handlers
	newCountryHandler(handler, uc, log)
	newTrendsHandler(handler, uc, log)
	newUtilityHandler(handler, log)

	return handler
}
