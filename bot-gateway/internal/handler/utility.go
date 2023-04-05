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

	// Handle start command
	handler.HandleMessage(h.start, th.CommandEqual(_cmdStart))
	// Handle help command
	handler.HandleMessage(h.help, th.CommandEqual(_cmdHelp))
	// Handle health check
	handler.HandleMessage(h.healthCheck, th.TextEqual("_Check"))
}

func (h *utilityHandler) start(bot *telego.Bot, message telego.Message) {
	id := message.Chat.ID

	err := response(bot, id, ui.InlineButton(_btnCountry), msg.UtilStart)
	if err != nil {
		h.log.Error().Err(err).
			Msg("Cannot send message")
	}
}

func (h *utilityHandler) help(bot *telego.Bot, message telego.Message) {
	id := message.Chat.ID

	err := response(bot, id, ui.InlineButtons(_btnSchedule, _btnCountry, _btnTrends), msg.UtilHelp)
	if err != nil {
		h.log.Error().Err(err).
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
		h.log.Error().Err(err).
			Msg("Cannot send message")
	}
}
