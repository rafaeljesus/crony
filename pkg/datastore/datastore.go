package datastore

import (
	"net/url"

	"github.com/jinzhu/gorm"
)

const (
	Postgres = "postgresql"
)

func New(dsn string) (*gorm.DB, error) {
	url, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}

	switch url.Scheme {
	case Postgres:
		c := PGConfig{
			Url:         dsn,
			MaxIdleConn: 10,
			MaxOpenConn: 100,
		}

		return NewPostgres(c)
	default:
		return nil, ErrUnknownDatabaseProvider
	}
}
