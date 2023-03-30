package grpc

import (
	"context"

	grpcscheduler "github.com/IamVladlen/trend-bot/proto/scheduler"
	"github.com/IamVladlen/trend-bot/scheduler-service/internal/usecase"
	"github.com/IamVladlen/trend-bot/scheduler-service/pkg/grpcsrv"
	"github.com/IamVladlen/trend-bot/scheduler-service/pkg/logger"
)

type handler struct {
	grpcscheduler.UnimplementedSchedulerServer

	uc  *usecase.UseCase
	log *logger.Logger
}

func New(uc *usecase.UseCase, srv *grpcsrv.Server, log *logger.Logger) {
	h := &handler{
		uc:  uc,
		log: log,
	}

	grpcscheduler.RegisterSchedulerServer(srv, h)
}

func (h *handler) GetScheduledMessages(ctx context.Context, req *grpcscheduler.GetScheduledRequest) (*grpcscheduler.GetScheduledResponse, error) {
	ids, err := h.uc.GetScheduledMessages(ctx, req.Interval)
	if err != nil {
		h.log.Error().Err(err).
			Msg("Cannot get scheduled messages")
	}

	return &grpcscheduler.GetScheduledResponse{Ids: ids}, err
}

func (h *handler) SetChatSchedule(ctx context.Context, req *grpcscheduler.SetChatRequest) (*grpcscheduler.SetChatResponse, error) {
	err := h.uc.SetChatSchedule(ctx, req.ChatId, req.Interval)
	if err != nil {
		h.log.Error().Err(err).
			Msg("Cannot set chat schedule")
	}

	return &grpcscheduler.SetChatResponse{}, err
}
