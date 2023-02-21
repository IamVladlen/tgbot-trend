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
	handler.HandleMessage(healthCheck, th.TextEqual("_Check"))

	// Initialize handlers
	newCountryHandler(handler, uc, log)
	newTrendsHandler(handler, uc, log)
	newUtilityHandler(handler)

	return handler
}

func healthCheck(bot *telego.Bot, message telego.Message) {
	m := tu.MessageWithEntities(
		tu.ID(message.Chat.ID),
		tu.Entity("I'm working, beep boop").Italic(),
	)
	bot.SendMessage(m)
}
