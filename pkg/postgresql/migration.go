package postgresql

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
)

const (
	sourceName = "go-bindata"
)

var (
	ErrDirtyMigration = errors.New("the migration is dirty")
)

// Migrate applies given database migration files for connection
func (c *connection) Migrate(version uint, files []string, afn func(name string) ([]byte, error)) (uint, error) {
	sourceInstance, err := bindata.WithInstance(bindata.Resource(files, afn))
	if err != nil {
		return 0, fmt.Errorf("preparing migrations source error: %w", err)
	}

	driver, err := postgres.WithInstance(c.DB.DB, &postgres.Config{})
	if err != nil {
		return 0, fmt.Errorf("preparing migrations driver error: %w", err)
	}

	m, err := migrate.NewWithInstance(sourceName, sourceInstance, driverName, driver)
	if err != nil {
		return 0, fmt.Errorf("preparing migrations instance error: %w", err)
	}

	if version == 0 {
		err = m.Up()
	} else {
		err = m.Migrate(version)
	}
	if err != nil {
		if err == migrate.ErrNoChange {
			err = nil
		} else {
			return 0, fmt.Errorf("migration error: %w", err)
		}
	}

	version, dirty, err := m.Version()
	if err != nil {
		return 0, fmt.Errorf("migration version error: %w", err)
	}
	if dirty {
		return 0, ErrDirtyMigration
	}

	return version, nil
}
