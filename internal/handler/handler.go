package handler

import (
	"github.com/IamVladlen/trend-bot/internal/usecase"
	"github.com/IamVladlen/trend-bot/pkg/logger"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func New(handler *th.BotHandler, uc *usecase.UseCase, log *logger.Logger) *th.BotHandler {
	// Health check
	handler.HandleMessage(healthCheck, healthCheckCond)
	newUtilityHandler(handler)

	// Initialize handlers
	newCountryHandler(handler, uc, log)
	newTrendsHandler(handler, uc, log)

	return handler
}

func healthCheck(bot *telego.Bot, message telego.Message) {
	bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "I'm working, beep boop"))
}
func healthCheckCond(update telego.Update) bool {
	return update.Message.Text == "_Check"
}
