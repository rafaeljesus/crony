package repos

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rafaeljesus/crony/pkg/models"
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
	e.CreatedAt = time.Now()
	return r.db.Create(e).Error
}

func (r *Event) FindById(id int) (e *models.Event, err error) {
	var event models.Event
	err = r.db.Find(&event, id).Error
	return &event, err
}

func (r *Event) Update(e *models.Event) error {
	e.UpdatedAt = time.Now()
	return r.db.Save(e).Error
}

func (r *Event) Delete(e *models.Event) error {
	e.DeletedAt = time.Now()
	return r.db.Delete(e).Error
}

func (r *Event) Search(q *models.Query) (events []models.Event, err error) {
	where := make(map[string]interface{})

	if !q.IsEmpty() {
		if q.Status != "" {
			where["status"] = q.Status
		}

		if q.Expression != "" {
			where["expression"] = q.Expression
		}
	}

	err = r.db.Limit(q.GetLimit()).Offset(q.Offset).Where(where).Find(&events).Error
	return events, err
}
