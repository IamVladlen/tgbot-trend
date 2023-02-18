package usecase

import (
	"github.com/IamVladlen/trend-bot/internal/entity"
	"github.com/IamVladlen/trend-bot/internal/repository"
	"github.com/IamVladlen/trend-bot/internal/webapi"
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

// New creates use case instance.
func New(repo *repository.Repository, api *webapi.WebAPI) *UseCase {
	return &UseCase{
		newCountryUC(repo.Country),
		newTrendsUC(repo.Country, api),
	}
}
