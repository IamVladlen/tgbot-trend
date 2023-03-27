package bot

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/IamVladlen/trend-bot/bot-gateway/config"
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/handler"
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/microservice"
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/repository"
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/usecase"
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/webapi"
	"github.com/IamVladlen/trend-bot/bot-gateway/pkg/logger"
	"github.com/IamVladlen/trend-bot/bot-gateway/pkg/mongodb"
	"github.com/IamVladlen/trend-bot/bot-gateway/pkg/redisdb"
	"github.com/IamVladlen/trend-bot/bot-gateway/pkg/tgbot"
	"github.com/IamVladlen/trend-bot/bot-gateway/pkg/ticker"
)

// Run starts the bot and connects all dependencies.
func Run(cfg *config.Config, log *logger.Logger) {
	// Databases
	mgdb := mongodb.New(mongodb.Deps{
		URI:      cfg.Mongo.URI,
		Username: cfg.Mongo.User,
		Password: cfg.Mongo.Password,
		DBName:   cfg.Mongo.DBName,
	})
	cache := redisdb.New(cfg.Redis.URI, cfg.Redis.Password)

	// Bot dependencies
	repo := repository.New(mgdb)
	service := microservice.New()
	webAPI := webapi.New(cache, log)
	uc := usecase.New(service, repo, webAPI)
	t := ticker.New()

	// Bot initialization and start
	bot := tgbot.New(cfg.Bot.Token)
	handler.New(handler.Deps{
		Bot:     bot.Bot,
		Handler: bot.Handler,
		UC:      uc,
		Log:     log,
		Ticker:  t,
	})

	t.StartAsync()
	bot.Start()

	// Graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	s := <-sigCh
	log.Info().Msg("Server is shutting down: " + s.String() + " signal")

	bot.Stop()

	if err := mgdb.Disconnect(); err != nil {
		log.Error().
			Err(err).
			Msg("error occurred while disconnecting from MongoDB")
	}

	cache.Close()
	t.Stop()
}
