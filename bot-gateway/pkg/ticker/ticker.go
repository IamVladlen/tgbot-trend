package ticker

import (
	"time"

	"github.com/go-co-op/gocron"
)

type Ticker struct {
	*gocron.Scheduler
}

// New creates ticker instance and starts ticker
// that handles scheduled functions. See Schedule
// for more info.
func New() *Ticker {
	s := gocron.NewScheduler(time.UTC)

	return &Ticker{
		s,
	}
}