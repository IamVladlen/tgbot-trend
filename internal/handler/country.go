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
	_cmdCountry = "Set country"
)

type countryHandler struct {
	uc  *usecase.UseCase
	log *logger.Logger

	isChangeable bool
}

func newCountryHandler(handler *th.BotHandler, uc *usecase.UseCase, log *logger.Logger) {
	h := &countryHandler{
		uc:  uc,
		log: log,

		isChangeable: false,
	}

	handler.HandleMessage(h.callChangeCountry, th.TextEqual(_cmdCountry))
	handler.HandleMessage(h.changeCountry, h.changeCountryCond)
}

// callChangeCountry puts the chat in waiting for the country id in the next message.
func (h *countryHandler) callChangeCountry(bot *telego.Bot, message telego.Message) {
	h.isChangeable = true

	kb := tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton("DE").WithText("🇩🇪"),
			tu.KeyboardButton("ES").WithText("🇪🇸"),
			tu.KeyboardButton("FR").WithText("🇫🇷"),
			tu.KeyboardButton("IT").WithText("🇮🇹"),
		),
		tu.KeyboardRow(
			tu.KeyboardButton("RU").WithText("🇬🇧"),
			tu.KeyboardButton("UA").WithText("🇷🇺"),
			tu.KeyboardButton("UK").WithText("🇺🇦"),
			tu.KeyboardButton("US").WithText("🇺🇸"),
		),
	)
	m := tu.Message(
		tu.ID(message.Chat.ID),
		msg.CallChangeCountry,
	).WithReplyMarkup(kb)

	bot.SendMessage(m)
}

// changeCountry changes country of fetched trends in chat.
func (h *countryHandler) changeCountry(bot *telego.Bot, message telego.Message) {
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

	kb := tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton("").WithText(_cmdCountry),
		),
		tu.KeyboardRow(
			tu.KeyboardButton("").WithText(_cmdTrends),
		),
	)
	m := tu.Message(
		tu.ID(message.Chat.ID),
		msg.ChangeCountrySucc,
	).WithReplyMarkup(kb)

	bot.SendMessage(m)
}

func (h *countryHandler) changeCountryCond(update telego.Update) bool {
	return h.isChangeable
}
