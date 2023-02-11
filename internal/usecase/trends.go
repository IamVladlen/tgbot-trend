package usecase

import (
	"fmt"

	"github.com/IamVladlen/trend-bot/internal/entity"
)

type TrendsUC struct {
	repo TrendsRepo
	api  TrendsWebAPI
}

// GetTrends sends struct with unmarshalled XML from
// Google API based on country set in the chat.
func (uc *TrendsUC) GetTrends(id int) (entity.Trends, error) {
	country, err := uc.repo.GetCountry(id)
	if err != nil {
		return entity.Trends{}, fmt.Errorf("usecase - GetTrends: %w", err)
	}

	trends, err := uc.api.GetTrends(country)
	if err != nil {
		return entity.Trends{}, fmt.Errorf("usecase - GetTrends: %w", err)
	}

	return trends, nil
}

func newTrendsUC(repo TrendsRepo, api TrendsWebAPI) *TrendsUC {
	return &TrendsUC{
		repo: repo,
		api:  api,
	}
}
