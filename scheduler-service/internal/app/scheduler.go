package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/IamVladlen/trend-bot/scheduler-service/config"
	"github.com/IamVladlen/trend-bot/scheduler-service/internal/handler/grpc"
	"github.com/IamVladlen/trend-bot/scheduler-service/internal/repository"
	"github.com/IamVladlen/trend-bot/scheduler-service/internal/usecase"
	"github.com/IamVladlen/trend-bot/scheduler-service/pkg/grpcsrv"
	"github.com/IamVladlen/trend-bot/scheduler-service/pkg/logger"
	"github.com/IamVladlen/trend-bot/scheduler-service/pkg/migration"
	"github.com/IamVladlen/trend-bot/scheduler-service/pkg/postgres"
)

func Run(log *logger.Logger, cfg *config.Config) {
	pg := postgres.New(cfg.PG.URL)
	migration.LoadMigrationSQL(cfg.PG.URL, cfg.PG.Migration)

	repo := repository.New(pg)

	uc := usecase.New(repo)

	grpcSrv := grpcsrv.New(cfg.GRPC.Port)
	grpc.New(uc, grpcSrv, log)

	log.Info().Msg("Scheduler service has successfully launched")

	// Graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	select {
	case s := <-sigCh:
		log.Info().
			Msg("Server is shutting down: " + s.String() + " signal")
	case err := <-grpcSrv.Notify():
		log.Error().Err(err).
			Msg("Server is shutting down due to error occurrence")
	}

	grpcSrv.Stop()
	pg.Close()
}
