package usecase

import (
	"github.com/IamVladlen/trend-bot/internal/entity"
	"github.com/IamVladlen/trend-bot/internal/repository"
	"github.com/IamVladlen/trend-bot/internal/webapi"
)

type UseCase struct {
	*ChatUC
	*TrendsUC
}

type Chat interface {
	ChangeCountry(chat entity.Chat) error
	GetCountry(id int) (string, error)
}

type TrendsRepo interface {
	GetCountry(id int) (string, error)
}

type TrendsWebAPI interface {
	GetTrends(country string) (entity.Trends, error)
}

func New(repo *repository.Repository, api *webapi.WebAPI) *UseCase {
	return &UseCase{
		newChatUC(repo.Chat),
		newTrendsUC(repo.Chat, api),
	}
}
