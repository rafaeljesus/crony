package repos

import (
	"github.com/EmpregoLigado/cron-srv/pkg/models"
	"github.com/jinzhu/gorm"
)

type EventRepo interface {
	Create(event *models.Event) (err error)
	FindById(id int) (event *models.Event, err error)
	Update(event *models.Event) (err error)
	Delete(event *models.Event) (err error)
	Search(query *models.Query) (events []models.Event, err error)
}

type Event struct {
	db *gorm.DB
}

func NewEvent(db *gorm.DB) *Event {
	return &Event{db}
}

func (r *Event) Create(e *models.Event) error {
	return r.db.Create(e).Error
}

func (r *Event) FindById(id int) (e *models.Event, err error) {
	err = r.db.Find(e, id).Error
	return
}

func (r *Event) Update(e *models.Event) error {
	return r.db.Save(e).Error
}

func (r *Event) Delete(e *models.Event) error {
	return r.db.Delete(e).Error
}

func (r *Event) Search(q *models.Query) (events []models.Event, err error) {
	if q.IsEmpty() {
		err = r.db.Find(&events).Error
		if err != nil {
			return
		}

		return
	}

	var db *gorm.DB
	if q.Status != "" {
		db = db.Where("status = ?", q.Status)
	}

	if q.Expression != "" {
		db = db.Where("expression = ?", q.Expression)
	}

	err = db.Find(events).Error

	return
}
