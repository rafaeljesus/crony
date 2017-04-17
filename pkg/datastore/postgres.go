package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rafaeljesus/crony/pkg/models"
)

type PGConfig struct {
	Url         string
	MaxIdleConn int
	MaxOpenConn int
	LogMode     bool
}

func NewPostgres(c PGConfig) (conn *gorm.DB, err error) {
	conn, err = gorm.Open("postgres", c.Url)
	if err != nil {
		return
	}

	if err = conn.DB().Ping(); err != nil {
		return
	}

	if c.MaxIdleConn == 0 {
		c.MaxIdleConn = 10
	}

	if c.MaxOpenConn == 0 {
		c.MaxOpenConn = 100
	}

	conn.DB().SetMaxIdleConns(c.MaxIdleConn)
	conn.DB().SetMaxOpenConns(c.MaxOpenConn)
	conn.LogMode(c.LogMode)

	// FIXME temporary
	conn.AutoMigrate(&models.Event{})

	return
}
