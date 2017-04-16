package checker

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type postgres struct {
	url string
}

func NewPostgres(url string) *postgres {
	return &postgres{url}
}

func (p *postgres) IsAlive() bool {
	conn, err := gorm.Open("postgres", p.url)
	if err != nil {
		return false
	}
	defer conn.Close()

	if err = conn.DB().Ping(); err != nil {
		return false
	}

	return true
}
