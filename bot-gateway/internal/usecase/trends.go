package usecase

import (
	"context"
	"fmt"

	"github.com/IamVladlen/trend-bot/bot-gateway/internal/entity"
)

type TrendsUC struct {
	service TrendsMicroservice
	repo    TrendsRepo
	api     TrendsWebAPI
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

func (uc *TrendsUC) GetScheduledMessages(ctx context.Context, interval string) ([]int64, error) {
	chatIds, err := uc.service.GetScheduledMessages(ctx, interval)
	if err != nil {
		return []int64{}, fmt.Errorf("usecase - GetScheduledMessage: %w", err)
	}

	return chatIds, nil
}

func (uc *TrendsUC) SetChatSchedule(ctx context.Context, chatId int64, interval string) error {
	if err := uc.service.SetChatSchedule(ctx, chatId, interval); err != nil {
		return fmt.Errorf("usecase - SetChatSchedule: %w", err)
	}

	return nil
}

func newTrendsUC(service TrendsMicroservice, repo TrendsRepo, api TrendsWebAPI) *TrendsUC {
	return &TrendsUC{
		service: service,
		repo:    repo,
		api:     api,
	}
}
