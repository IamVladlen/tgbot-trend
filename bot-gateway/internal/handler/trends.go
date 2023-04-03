package handler

import (
	"context"
	"time"

	"github.com/IamVladlen/trend-bot/bot-gateway/internal/handler/msg"
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/handler/ui"
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/usecase"
	"github.com/IamVladlen/trend-bot/bot-gateway/pkg/logger"
	"github.com/IamVladlen/trend-bot/bot-gateway/pkg/ticker"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

const (
	_btnTrends   = "Get trends"
	_btnDaily    = "Daily"
	_btnWeekly   = "Weekly"
	_btnUnsubscribe    = "Unsubscribe"
	_btnSchedule = "Newsletter"
)

type trendsHandler struct {
	uc  *usecase.UseCase
	log *logger.Logger
	bot *telego.Bot
}

func newTrendsHandler(bot *telego.Bot, handler *th.BotHandler, uc *usecase.UseCase, log *logger.Logger, t *ticker.Ticker) {
	h := &trendsHandler{
		uc:  uc,
		log: log,
		bot: bot,
	}

	// Handle scheduled messages everyday
	_, err := t.Every(1).Day().At("20:00").Do(func() { h.getScheduledMessages(bot, "Daily") })
	if err != nil {
		log.Error().Err(err).
			Msg("Cannot send message")
	}

	// Handle scheduled messages everyday
	_, err = t.Every(1).Minute().Do(func() { h.getScheduledMessages(bot, "Daily") })
	if err != nil {
		log.Error().Err(err).
			Msg("Cannot send message")
	}

	// Handle scheduled messages every week
	_, err = t.Every(1).Week().Weekday(time.Sunday).At("20:00").Do(func() { h.getScheduledMessages(bot, "Weekly") })
	if err != nil {
		log.Error().Err(err).
			Msg("Cannot send message")
	}

	// Handle manual trends getting
	handler.HandleCallbackQuery(h.getTrends, th.CallbackDataEqual(_btnTrends))
	// Handle chat schedule trigger
	handler.HandleCallbackQuery(h.callSetChatSchedule, th.CallbackDataEqual(_btnSchedule))
	// Handle chat schedule
	handler.HandleCallbackQuery(h.setChatSchedule, th.CallbackDataEqual(_btnDaily))
	handler.HandleCallbackQuery(h.setChatSchedule, th.CallbackDataEqual(_btnWeekly))
	handler.HandleCallbackQuery(h.setChatSchedule, th.CallbackDataEqual(_btnUnsubscribe))
}

// getTrends sends a list of trends.
func (h *trendsHandler) getTrends(bot *telego.Bot, query telego.CallbackQuery) {
	id := query.Message.Chat.ID

	trends, err := h.uc.GetTrends(int(id))
	if err != nil {
		h.log.Error().Err(err).
			Msg("Cannot get trends")

		err := response(bot, id, ui.InlineButtons(_btnSchedule, _btnCountry, _btnTrends), msg.GetTrendsFailFetch)
		if err != nil {
			h.log.Error().Err(err).
				Msg("Cannot send message")
		}

		return
	}

	text := trends.EntityString()
	err = responseEntities(bot, id, ui.InlineButtons(_btnSchedule, _btnCountry, _btnTrends), text...)
	if err != nil {
		h.log.Error().Err(err).
			Msg("Cannot send message")
	}
}

// getScheduledMessages is a callback for cron scheduler and a wrapper
// over getTrends function that cycles through received chat ids.
func (h *trendsHandler) getScheduledMessages(bot *telego.Bot, interval string) {
	chatIds, err := h.uc.GetScheduledMessages(context.Background(), interval)
	if err != nil {
		h.log.Error().Err(err).
			Msg("Cannot get trends")
	}

	for _, id := range chatIds {
		trends, err := h.uc.GetTrends(int(id))
		if err != nil {
			h.log.Error().Err(err).
				Msg("Cannot get trends")

			err := response(bot, id, ui.InlineButtons(_btnSchedule, _btnCountry, _btnTrends), msg.GetTrendsFailFetch)
			if err != nil {
				h.log.Error().Err(err).
					Msg("Cannot send message")
			}

			return
		}

		text := trends.EntityString()
		err = responseEntities(bot, id, ui.InlineButtons(_btnSchedule, _btnCountry, _btnTrends), text...)
		if err != nil {
			h.log.Error().Err(err).
				Msg("Cannot send message")
		}
	}
}

// callSetChatSchedule puts the chat into a pending interval state with the next message.
func (h *trendsHandler) callSetChatSchedule(bot *telego.Bot, query telego.CallbackQuery) {
	id := query.Message.Chat.ID

	err := response(bot, id, ui.InlineButtonsSchedule(_btnDaily, _btnWeekly, _btnUnsubscribe), msg.CallSetChatSchedule)
	if err != nil {
		h.log.Error().Err(err).
			Msg("Cannot send message")
	}
}

// TODO: Delete previous message if invoked

// setChatSchedule sets schedule for newsletter.
func (h *trendsHandler) setChatSchedule(bot *telego.Bot, query telego.CallbackQuery) {
	id := query.Message.Chat.ID
	interval := query.Data

	if err := h.uc.SetChatSchedule(context.Background(), id, interval); err != nil {
		h.log.Error().Err(err).
			Msg("Cannot set chat schedule")

		err := response(bot, id, ui.InlineButtonsSchedule(_btnSchedule, _btnCountry, _btnTrends), msg.SetChatScheduleErr)
		if err != nil {
			h.log.Error().Err(err).
				Msg("Cannot send message")
		}

		return
	}

	err := response(bot, id, ui.InlineButtonsSchedule(_btnSchedule, _btnCountry, _btnTrends), msg.SetChatScheduleSucc)
	if err != nil {
		h.log.Error().Err(err).
			Msg("Cannot send message")
	}
}
