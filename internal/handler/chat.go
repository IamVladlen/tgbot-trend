package handler

import (
	"github.com/IamVladlen/trend-bot/internal/entity"
	"github.com/IamVladlen/trend-bot/internal/handler/msg"
	"github.com/IamVladlen/trend-bot/internal/usecase"
	"github.com/IamVladlen/trend-bot/pkg/logger"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

const (
	_cmdCountry = "country"
)

type chatHandler struct {
	uc  *usecase.UseCase
	log *logger.Logger

	isChangeable bool
}

func newChatHandler(handler *th.BotHandler, uc *usecase.UseCase, log *logger.Logger) {
	h := &chatHandler{
		uc:  uc,
		log: log,

		isChangeable: false,
	}

	handler.HandleMessage(h.callChangeCountry, th.CommandEqual(_cmdCountry))
	handler.HandleMessage(h.changeCountry, h.changeCountryCond)
}

// callChangeCountry puts the chat in waiting for the country id in the next message.
func (h *chatHandler) callChangeCountry(bot *telego.Bot, message telego.Message) {
	h.isChangeable = true
	m := tu.Message(
		tu.ID(message.Chat.ID),
		msg.CallChangeCountry,
	)

	bot.SendMessage(m)
}

// changeCountry changes country of fetched trends in chat.
func (h *chatHandler) changeCountry(bot *telego.Bot, message telego.Message) {
	chat := entity.Chat{
		ChatId:  int(message.Chat.ID),
		Country: message.Text,
	}

	if err := h.uc.ChangeCountry(chat); err != nil {
		h.isChangeable = false
		h.log.Error().Msg("can't change country: " + err.Error())
		m := tu.Message(
			tu.ID(message.Chat.ID),
			msg.ChangeCountryFail,
		)

		bot.SendMessage(m)
		return
	}

	h.isChangeable = false
	m := tu.Message(
		tu.ID(message.Chat.ID),
		msg.ChangeCountrySucc,
	)

	bot.SendMessage(m)
}

func (h *chatHandler) changeCountryCond(update telego.Update) bool {
	return h.isChangeable
}
