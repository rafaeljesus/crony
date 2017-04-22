package scheduler

import (
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/rafaeljesus/crony/pkg/models"
	"github.com/rafaeljesus/crony/pkg/repos"
	"github.com/rafaeljesus/crony/pkg/runner"
	"github.com/robfig/cron"
)

type Scheduler interface {
	Create(event *models.Event) error
	Update(event *models.Event) error
	Delete(id uint) error
	Find(id uint) (*cron.Cron, error)
	ScheduleAll(r repos.EventRepo)
}

type scheduler struct {
	sync.RWMutex
	Kv   map[uint]*cron.Cron
	Cron *cron.Cron
}

func New() Scheduler {
	s := &scheduler{
		Kv:   make(map[uint]*cron.Cron),
		Cron: cron.New(),
	}

	s.Cron.Start()

	return s
}

func (s *scheduler) Create(e *models.Event) (err error) {
	s.Cron.AddFunc(e.Expression, func() {
		c := &runner.Config{
			Url:     e.Url,
			Retries: e.Retries,
			Timeout: e.RetryTimeout,
		}

		r := runner.New()
		r.Run() <- c
	})

	s.Lock()
	defer s.Unlock()

	s.Kv[e.Id] = s.Cron

	return
}

func (s *scheduler) Find(id uint) (cron *cron.Cron, err error) {
	s.Lock()
	defer s.Unlock()

	cron, found := s.Kv[id]
	if !found {
		err = ErrEventNotExist
		return
	}

	return
}

func (s *scheduler) Update(e *models.Event) (err error) {
	if err = s.Delete(e.Id); err != nil {
		return
	}

	return s.Create(e)
}

func (s *scheduler) Delete(id uint) (err error) {
	s.Lock()
	defer s.Unlock()

	_, found := s.Kv[id]
	if !found {
		err = ErrEventNotExist
		return
	}

	s.Kv[id].Stop()
	s.Kv[id] = nil

	return
}

func (s *scheduler) ScheduleAll(r repos.EventRepo) {
	events, err := r.Search(&models.Query{})
	if err != nil {
		log.Error("Failed to find events!")
		return
	}

	for _, e := range events {
		if err = s.Create(&e); err != nil {
			log.Error("Failed to create event!")
		}
	}
}
