package usecase

import (
	"context"

	"github.com/IamVladlen/trend-bot/bot-gateway/internal/entity"
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/microservice"
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/repository"
	"github.com/IamVladlen/trend-bot/bot-gateway/internal/webapi"
)

//go:generate mockgen -source=usecase.go -destination=./mocks/mocks.go -package=mocks

type UseCase struct {
	*CountryUC
	*TrendsUC
}

type CountryRepo interface {
	ChangeCountry(id int, country string) error
	GetCountry(id int) (string, error)
}

type TrendsRepo interface {
	GetCountry(id int) (string, error)
}

type TrendsWebAPI interface {
	GetTrends(country string) (entity.Trends, error)
}

type TrendsMicroservice interface {
	GetScheduledMessages(ctx context.Context, interval string) ([]int64, error)
	SetChatSchedule(ctx context.Context, chatId int64, interval string) error
}

// New creates use case instance.
func New(service *microservice.Microservice, repo *repository.Repository, api *webapi.WebAPI) *UseCase {
	return &UseCase{
		newCountryUC(repo.Country),
		newTrendsUC(service, repo.Country, api),
	}
}
