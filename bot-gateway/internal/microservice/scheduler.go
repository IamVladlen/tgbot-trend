package microservice

import (
	"context"
	"fmt"

	grpcscheduler "github.com/IamVladlen/tgbot-trend/proto/scheduler"
)

type scheduler struct {
	grpc grpcscheduler.SchedulerClient
}

func (s *scheduler) GetScheduledMessages(ctx context.Context, interval string) ([]int64, error) {
	req := &grpcscheduler.GetScheduledRequest{
		Interval: interval,
	}

	res, err := s.grpc.GetScheduledMessages(ctx, req)
	if err != nil {
		return []int64{}, fmt.Errorf("microservice - GetScheduledMessages: %w", err)
	}

	return res.GetIds(), nil
}

func (s *scheduler) SetChatSchedule(ctx context.Context, chatId int64, interval string) error {
	req := &grpcscheduler.SetChatRequest{
		ChatId:   chatId,
		Interval: interval,
	}

	_, err := s.grpc.SetChatSchedule(ctx, req)
	if err != nil {
		return fmt.Errorf("microservice - SetChatSchedule: %w", err)
	}

	return nil
}

func newSchedulerService(grpcclt grpcscheduler.SchedulerClient) *scheduler {
	return &scheduler{
		grpcclt,
	}
}
