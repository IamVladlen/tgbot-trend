package handler

import (
	"github.com/IamVladlen/trend-bot/internal/handler/msg"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

const (
	_cmdStart = "start"
	_cmdHelp  = "help"
)

type utilityHandler struct{}

func newUtilityHandler(handler *th.BotHandler) {
	h := &utilityHandler{}

	handler.HandleMessage(h.start, th.CommandEqual(_cmdStart))
	handler.HandleMessage(h.help, th.CommandEqual(_cmdHelp))
}

func (h *utilityHandler) start(bot *telego.Bot, message telego.Message) {
	kb := tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton("/country").WithText("/"+_cmdCountry),
		),
	)
	
	m := tu.Message(
		tu.ID(message.Chat.ID),
		msg.UtilStart,
	).WithReplyMarkup(kb)

	bot.SendMessage(m)
}

func (h *utilityHandler) help(bot *telego.Bot, message telego.Message) {
	kb := tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton("/country").WithText("/"+_cmdCountry),
		),
		tu.KeyboardRow(
			tu.KeyboardButton("/trends").WithText("/"+_cmdTrends),
		),
	)
	m := tu.Message(
		tu.ID(message.Chat.ID),
		msg.UtilHelp,
	).WithReplyMarkup(kb)

	bot.SendMessage(m)
}
