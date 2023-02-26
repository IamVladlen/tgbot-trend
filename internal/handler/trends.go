package handler

import (
	"github.com/IamVladlen/trend-bot/internal/handler/msg"
	"github.com/IamVladlen/trend-bot/internal/handler/ui"
	"github.com/IamVladlen/trend-bot/internal/usecase"
	"github.com/IamVladlen/trend-bot/pkg/logger"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

const (
	_cmdTrends = "Get trends"
)

type trendsHandler struct {
	uc  *usecase.UseCase
	log *logger.Logger
}

func newTrendsHandler(handler *th.BotHandler, uc *usecase.UseCase, log *logger.Logger) {
	h := &trendsHandler{
		uc:  uc,
		log: log,
	}

	handler.HandleCallbackQuery(h.getTrends, th.CallbackDataEqual(_cmdTrends))
}

// getTrends sends a list of trends.
func (h *trendsHandler) getTrends(bot *telego.Bot, query telego.CallbackQuery) {
	id := query.Message.Chat.ID

	trends, err := h.uc.GetTrends(int(id))
	if err != nil {
		h.log.Error().
			Err(err).
			Msg("can't get trends")

		m := tu.Message(
			tu.ID(id),
			msg.GetTrendsFailFetch,
		).WithReplyMarkup(ui.InlineButtons(_cmdCountry, _cmdTrends))
		bot.SendMessage(m)

		return
	}

	text := trends.EntityString()
	m := tu.MessageWithEntities(
		tu.ID(id),
		text...,
	).WithReplyMarkup(ui.InlineButtons(_cmdCountry, _cmdTrends))

	bot.SendMessage(m)
}
