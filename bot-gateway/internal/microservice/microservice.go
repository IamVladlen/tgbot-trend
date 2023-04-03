package microservice

import grpcscheduler "github.com/IamVladlen/tgbot-trend/proto/scheduler"

type Microservice struct {
	*scheduler
}

func New(grpcclt grpcscheduler.SchedulerClient) *Microservice {
	return &Microservice{
		newSchedulerService(grpcclt),
	}
}
