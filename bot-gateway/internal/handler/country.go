package handler

import (
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/handler/msg"
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/handler/ui"
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/usecase"
	"github.com/IamVladlen/trend-bot/bot-gateway/pkg/logger"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

const (
	_btnCountry = "Set country"
)

type countryHandler struct {
	uc  *usecase.UseCase
	log *logger.Logger
}

func newCountryHandler(handler *th.BotHandler, uc *usecase.UseCase, log *logger.Logger) {
	h := &countryHandler{
		uc:  uc,
		log: log,
	}

	// Handle country changing trigger
	handler.HandleCallbackQuery(h.callChangeCountry, th.CallbackDataEqual(_btnCountry))
	// Handle country changing
	handler.HandleCallbackQuery(h.changeCountry, h.callChangeCountryCond)
}

// callChangeCountry puts the chat in waiting for the country id in the next message.
func (h *countryHandler) callChangeCountry(bot *telego.Bot, query telego.CallbackQuery) {
	id := query.Message.Chat.ID
	countries := []string{"🇩🇪", "🇪🇸", "🇫🇷", "🇮🇹", "🇬🇧", "🇷🇺", "🇺🇦", "🇺🇸"}

	err := response(bot, id, ui.InlineCountries(countries), msg.CallChangeCountry)
	if err != nil {
		h.log.Error().Err(err).
			Msg("Cannot send message")
	}
}

func (h *countryHandler) callChangeCountryCond(update telego.Update) bool {
	str, err := validateCountry(update.CallbackQuery.Data)
	if err != nil {
		return false
	}
	return str != ""
}

// changeCountry changes country of fetched trends in chat.
func (h *countryHandler) changeCountry(bot *telego.Bot, query telego.CallbackQuery) {
	id := query.Message.Chat.ID
	country, err := validateCountry(query.Data)
	if err != nil {
		h.log.Error().Err(err).
			Msg("can't change country")

		err := response(bot, id, ui.InlineButton(_btnCountry), msg.ChangeCountryInputFail)
		if err != nil {
			h.log.Error().Err(err).
				Msg("Cannot send message")
		}

		return
	}

	if err := h.uc.ChangeCountry(int(id), country); err != nil {
		h.log.Error().Err(err).
			Msg("can't change country")

		err := response(bot, id, ui.InlineButton(_btnCountry), msg.ChangeCountryServerFail)
		if err != nil {
			h.log.Error().Err(err).
				Msg("Cannot send message")
		}

		return
	}

	err = bot.DeleteMessage(&telego.DeleteMessageParams{ChatID: tu.ID(id), MessageID: query.Message.MessageID})
	if err != nil {
		h.log.Error().Err(err).
			Msg("Cannot send message")
	}

	response(bot, id, ui.InlineButtons(_btnSchedule, _btnCountry, _btnTrends), msg.ChangeCountrySucc(query.Data))
	if err != nil {
		h.log.Error().Err(err).
			Msg("Cannot send message")
	}
}

// TODO: Switch to map after increasing the number of countries

// validateCountry converts emoji to plain text and returns
// an error if there is no reference.
func validateCountry(text string) (string, error) {
	switch text {
	case "🇩🇪":
		return "DE", nil
	case "🇪🇸":
		return "ES", nil
	case "🇫🇷":
		return "FR", nil
	case "🇮🇹":
		return "IT", nil
	case "🇬🇧":
		return "GB", nil
	case "🇷🇺":
		return "RU", nil
	case "🇺🇦":
		return "UA", nil
	case "🇺🇸", "🇺🇲":
		return "US", nil
	default:
		return "", errInvalidCountry
	}
}
