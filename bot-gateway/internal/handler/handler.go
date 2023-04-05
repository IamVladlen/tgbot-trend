package handler

import (
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/usecase"
	"github.com/IamVladlen/trend-bot/bot-gateway/pkg/logger"
	"github.com/IamVladlen/trend-bot/bot-gateway/pkg/ticker"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type Deps struct {
	Bot     *telego.Bot
	Handler *th.BotHandler
	UC      *usecase.UseCase
	Log     *logger.Logger
	Ticker  *ticker.Ticker
}

func New(deps Deps) {
	// Initialize handlers
	newCountryHandler(deps.Handler, deps.UC, deps.Log)
	newTrendsHandler(deps.Bot, deps.Handler, deps.UC, deps.Log, deps.Ticker)
	newUtilityHandler(deps.Handler, deps.Log)
}
