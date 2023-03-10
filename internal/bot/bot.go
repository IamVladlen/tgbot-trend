package bot

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/IamVladlen/trend-bot/config"
	"github.com/IamVladlen/trend-bot/internal/handler"
	"github.com/IamVladlen/trend-bot/internal/repository"
	"github.com/IamVladlen/trend-bot/internal/usecase"
	"github.com/IamVladlen/trend-bot/internal/webapi"
	"github.com/IamVladlen/trend-bot/pkg/logger"
	"github.com/IamVladlen/trend-bot/pkg/mongodb"
	"github.com/IamVladlen/trend-bot/pkg/redisdb"
	"github.com/IamVladlen/trend-bot/pkg/tgbot"
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
	webAPI := webapi.New(cache, log)
	uc := usecase.New(repo, webAPI)

	// Bot initialization and start
	bot := tgbot.New(cfg.Bot.Token)
	handler.New(bot.Handler, uc, log)

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
}
