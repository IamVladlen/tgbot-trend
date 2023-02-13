package handler

import (
	"github.com/IamVladlen/trend-bot/internal/entity"
	"github.com/IamVladlen/trend-bot/internal/handler/msg"
	"github.com/IamVladlen/trend-bot/internal/handler/ui"
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

	handler.HandleCallbackQuery(h.callChangeCountry, th.CallbackDataEqual(_cmdCountry))
	handler.HandleCallbackQuery(h.changeCountry, th.AnyCallbackQuery(), h.changeCountryCond)
}

// callChangeCountry puts the chat in waiting for the country id in the next message.
func (h *countryHandler) callChangeCountry(bot *telego.Bot, query telego.CallbackQuery) {
	id := query.Message.Chat.ID
	h.isChangeable = true
	countries := []string{"🇩🇪", "🇪🇸", "🇫🇷", "🇮🇹", "🇬🇧", "🇷🇺", "🇺🇦", "🇺🇸"}

	m := tu.Message(
		tu.ID(id),
		msg.CallChangeCountry,
	).WithReplyMarkup(ui.InlineCountries(countries))

	bot.SendMessage(m)
}

// changeCountry changes country of fetched trends in chat.
func (h *countryHandler) changeCountry(bot *telego.Bot, query telego.CallbackQuery) {
	id := query.Message.Chat.ID
	country := query.Data

	chat := entity.Chat{
		ChatId:  int(id),
		Country: country,
	}

	if err := h.uc.ChangeCountry(chat); err != nil {
		h.isChangeable = false
		h.log.Error().Msg("can't change country: " + err.Error())

		m := tu.Message(
			tu.ID(id),
			msg.ChangeCountryFail,
		).WithReplyMarkup(ui.InlineButton(_cmdCountry))
		bot.SendMessage(m)

		return
	}

	h.isChangeable = false
	
	bot.DeleteMessage(&telego.DeleteMessageParams{ChatID: tu.ID(id), MessageID: query.Message.MessageID})

	m := tu.Message(
		tu.ID(id),
		msg.ChangeCountrySucc(country),
	).WithReplyMarkup(ui.InlineButtons(_cmdCountry, _cmdTrends))
	bot.SendMessage(m)
}

func (h *countryHandler) changeCountryCond(update telego.Update) bool {
	return h.isChangeable
}
