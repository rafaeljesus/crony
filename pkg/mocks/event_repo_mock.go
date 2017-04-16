package mocks

import (
	"github.com/EmpregoLigado/cron-srv/pkg/models"
)

type EventRepoMock struct {
	Created      bool
	Updated      bool
	Deleted      bool
	Found        bool
	Searched     bool
	ByStatus     bool
	ByExpression bool
}

func NewEventRepo() *EventRepoMock {
	return &EventRepoMock{
		Created:      false,
		Updated:      false,
		Deleted:      false,
		Found:        false,
		Searched:     false,
		ByStatus:     false,
		ByExpression: false,
	}
}

func (repo *EventRepoMock) Create(event *models.Event) (err error) {
	repo.Created = true
	return
}

func (repo *EventRepoMock) FindById(id int) (event *models.Event, err error) {
	repo.Found = true
	event = &models.Event{Id: 1}
	return
}

func (repo *EventRepoMock) Update(event *models.Event) (err error) {
	repo.Updated = true
	return
}

func (repo *EventRepoMock) Delete(event *models.Event) (err error) {
	repo.Deleted = true
	return
}

func (repo *EventRepoMock) Search(sc *models.Query) (events []models.Event, err error) {
	events = append(events, models.Event{Expression: "* * * * * *"})

	switch true {
	case sc.Status != "":
		repo.ByStatus = true
	case sc.Expression != "":
		repo.ByExpression = true
	default:
		repo.Searched = true
	}

	return
}
