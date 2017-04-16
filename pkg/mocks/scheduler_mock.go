package mocks

import (
	"github.com/EmpregoLigado/cron-srv/pkg/models"
	"github.com/EmpregoLigado/cron-srv/pkg/repos"
	"github.com/robfig/cron"
)

type SchedulerMock struct {
	Created bool
	Updated bool
	Deleted bool
}

func NewScheduler() *SchedulerMock {
	return &SchedulerMock{
		Created: false,
		Updated: false,
		Deleted: false,
	}
}

func (s *SchedulerMock) Create(event *models.Event) (err error) {
	s.Created = true
	return
}

func (s *SchedulerMock) Update(event *models.Event) (err error) {
	s.Updated = true
	return
}

func (s *SchedulerMock) Delete(id uint) (err error) {
	s.Deleted = true
	return
}

func (s SchedulerMock) Find(id uint) (c *cron.Cron, err error) {
	return
}

func (s *SchedulerMock) ScheduleAll(repo repos.EventRepo) {
	return
}
