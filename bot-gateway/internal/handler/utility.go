package handler

import (
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/handler/msg"
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/handler/ui"
	"github.com/IamVladlen/trend-bot/bot-gateway/pkg/logger"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

const (
	_cmdStart = "start"
	_cmdHelp  = "help"
)

type utilityHandler struct {
	log *logger.Logger
}

func newUtilityHandler(handler *th.BotHandler, log *logger.Logger) {
	h := &utilityHandler{
		log: log,
	}

	handler.HandleMessage(h.start, th.CommandEqual(_cmdStart))
	handler.HandleMessage(h.help, th.CommandEqual(_cmdHelp))
	handler.HandleMessage(h.healthCheck, th.TextEqual("_Check"))
}

func (h *utilityHandler) start(bot *telego.Bot, message telego.Message) {
	m := tu.Message(
		tu.ID(message.Chat.ID),
		msg.UtilStart,
	).WithReplyMarkup(ui.InlineButton(_cmdCountry))

	_, err := bot.SendMessage(m)
	if err != nil {
		h.log.Error().
			Err(err).
			Msg("Cannot send message")
	}
}

func (h *utilityHandler) help(bot *telego.Bot, message telego.Message) {
	m := tu.Message(
		tu.ID(message.Chat.ID),
		msg.UtilHelp,
	).WithReplyMarkup(ui.InlineButtons(_cmdCountry, _cmdTrends))

	_, err := bot.SendMessage(m)
	if err != nil {
		h.log.Error().
			Err(err).
			Msg("Cannot send message")
	}
}

func (h *utilityHandler) healthCheck(bot *telego.Bot, message telego.Message) {
	m := tu.MessageWithEntities(
		tu.ID(message.Chat.ID),
		tu.Entity("I'm working, beep boop").Italic(),
	)

	_, err := bot.SendMessage(m)
	if err != nil {
		h.log.Error().
			Err(err).
			Msg("Cannot send message")
	}
}
