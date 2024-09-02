package internals

import (
	"time"

	"github.com/go-co-op/gocron"
)

type Scheduler struct {
	scheduler *gocron.Scheduler
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		scheduler: gocron.NewScheduler(time.UTC),
	}
}

func (s *Scheduler) GetScheduler() *gocron.Scheduler {
	return s.scheduler
}
