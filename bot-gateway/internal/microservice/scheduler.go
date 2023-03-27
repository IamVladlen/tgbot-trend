package microservice

import "context"

type scheduler struct {
}

func (s *scheduler) GetScheduledMessages(ctx context.Context, interval string) ([]int64, error) {
	return []int64{}, nil
}

func (s *scheduler) SetChatSchedule(ctx context.Context, chatId int64, interval string) error {
	return nil
}

func newSchedulerService() *scheduler {
	return &scheduler{}
}