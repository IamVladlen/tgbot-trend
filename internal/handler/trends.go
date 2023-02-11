package handler

import (
	"github.com/IamVladlen/trend-bot/internal/handler/msg"
	"github.com/IamVladlen/trend-bot/internal/usecase"
	"github.com/IamVladlen/trend-bot/pkg/logger"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

const (
	_cmdTrends = "trends"
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

	handler.HandleMessage(h.getTrends, th.CommandEqual(_cmdTrends))
}

// getTrends sends a list of trends.
func (h *trendsHandler) getTrends(bot *telego.Bot, message telego.Message) {
	trends, err := h.uc.GetTrends(int(message.Chat.ID))
	if err != nil {
		h.log.Error().Msg("can't get trends: " + err.Error())
		m := tu.Message(
			tu.ID(message.Chat.ID),
			msg.GetTrendsFailFetch,
		)

		bot.SendMessage(m)
		return
	}

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
		trends.String(),
	).WithReplyMarkup(kb)

	bot.SendMessage(m)
}
