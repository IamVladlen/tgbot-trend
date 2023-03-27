package microservice

type Microservice struct {
	*scheduler
}

func New() *Microservice {
	return &Microservice{
		newSchedulerService(),
	}
}
