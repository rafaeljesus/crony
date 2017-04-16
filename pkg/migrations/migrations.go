package migrations

import (
	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/mattes/migrate/migrate"
)

type Migrations struct {
	url  string
	path string
}

func New(url, path string) *Migrations {
	return &Migrations{url, path}
}

func (m *Migrations) Up() (errs []error, ok bool) {
	errs, ok = migrate.UpSync(m.url, m.path)

	return
}
