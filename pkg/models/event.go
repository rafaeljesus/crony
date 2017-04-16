package models

import (
	"time"
)

type Event struct {
	Id         uint       `json:"id",sql:"primary_key"`
	Url        string     `json:"url",sql:"not null"`
	Expression string     `json:"expression",sql:"not null"`
	Status     string     `json:"status",sql:"not null"`
	Retries    int64      `json:"retries`
	Timeout    int64      `json:"timeout`
	CreatedAt  time.Time  `json:"created_at",sql:"not null"`
	UpdatedAt  time.Time  `json:"updated_at",sql:"not null"`
	DeletedAt  *time.Time `json:"created_at,omitempty"`
}
