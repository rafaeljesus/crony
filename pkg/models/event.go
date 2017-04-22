package models

import (
	"regexp"
	"time"
)

var (
	Active   = "active"
	Inactive = "inactive"
)

type Event struct {
	Id           uint          `json:"id" gorm:"primary_key"`
	Url          string        `json:"url" gorm:"not null"`
	Expression   string        `json:"expression" gorm:"not null"`
	Status       string        `json:"status" gorm:"index:idx_status" gorm:"not null"`
	Retries      int64         `json:"retries"`
	RetryTimeout time.Duration `json:"retry_timeout"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DeletedAt    time.Time     `json:"deleted_at,omitempty"`
}

func NewEvent() *Event {
	return &Event{
		Status:       Active,
		Retries:      1,
		RetryTimeout: time.Second * 1,
	}
}

func (e *Event) Validate() (errors map[string]string, ok bool) {
	errors = make(map[string]string)
	if e.Url == "" {
		errors["url"] = "field url is mandatory"
	}

	if e.Expression == "" {
		errors["expression"] = "field expression is mandatory"
	}

	match, err := regexp.MatchString("^active$|^inactive$", e.Status)
	if err != nil || !match {
		errors["status"] = "field status must be active or inactive"
	}

	if e.Retries < 0 || e.Retries > 10 {
		errors["retries"] = "field retries must be between 0 and 10"
	}

	ok = len(errors) == 0

	return
}

func (e *Event) SetAttributes(newEvent *Event) {
	if newEvent.Url != "" {
		e.Url = newEvent.Url
	}

	if newEvent.Expression != "" {
		e.Expression = newEvent.Expression
	}

	if newEvent.Status != "" {
		e.Status = newEvent.Status
	}

	if newEvent.Retries > 0 {
		e.Retries = newEvent.Retries
	}

	if newEvent.RetryTimeout > 0 {
		e.RetryTimeout = newEvent.RetryTimeout
	}
}
