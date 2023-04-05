package usecase

import (
	"context"
)

type UseCase struct {
	repo UseCaseRepo
}

type UseCaseRepo interface {
	GetScheduledMessages(ctx context.Context, interval string) ([]int64, error)
	SetChatSchedule(ctx context.Context, chatId int64, interval string) error
}

func (uc *UseCase) GetScheduledMessages(ctx context.Context, interval string) ([]int64, error) {
	return uc.repo.GetScheduledMessages(ctx, interval)
}

func (uc *UseCase) SetChatSchedule(ctx context.Context, chatId int64, interval string) error {
	return uc.repo.SetChatSchedule(ctx, chatId, interval)
}

func New(repo UseCaseRepo) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
