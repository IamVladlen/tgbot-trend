package bot

import (
	"github.com/IamVladlen/trend-bot/config"
	"github.com/IamVladlen/trend-bot/internal/handler"
	"github.com/IamVladlen/trend-bot/internal/repository"
	"github.com/IamVladlen/trend-bot/internal/usecase"
	"github.com/IamVladlen/trend-bot/internal/webapi"
	"github.com/IamVladlen/trend-bot/pkg/logger"
	"github.com/IamVladlen/trend-bot/pkg/mongodb"
	"github.com/IamVladlen/trend-bot/pkg/redisdb"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

// TODO: Move bot initialization to pkg
// TODO: Implement graceful shutdown

// Run starts the bot and connects all dependencies
func Run(cfg *config.Config, log *logger.Logger) {
	mgdb := mongodb.New(mongodb.Deps{
		URI:      cfg.Mongo.URI,
		Username: cfg.Mongo.User,
		Password: cfg.Mongo.Password,
		DBName:   cfg.Mongo.DBName,
	})
	repo := repository.New(mgdb)
	cache := redisdb.New(cfg.Redis.URI, cfg.Redis.Password)
	webAPI := webapi.New(cache, log)

	uc := usecase.New(repo, webAPI)

	bot, err := telego.NewBot(cfg.Bot.Token)
	if err != nil {
		log.Fatal().Msgf("app - Run - telego.NewBot: %s", err.Error())
	}

	updates, err := bot.UpdatesViaLongPolling(nil)
	if err != nil {
		log.Fatal().Msgf("app - Run - telego.NewBot: %s", err.Error())
	}
	defer bot.StopLongPolling()

	h, err := th.NewBotHandler(bot, updates)
	if err != nil {
		log.Fatal().Msgf("app - Run - th.NewBotHandler: %s", err.Error())
	}
	defer h.Stop()

	handler.New(h, uc, log)

	log.Info().Msg("Bot has successfully launched")
	h.Start()
}
