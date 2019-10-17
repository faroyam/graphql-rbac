package migrations

import (
	"graphql-rbac/pkg/logger"
	"graphql-rbac/pkg/postgresql"

	"go.uber.org/zap"
)

//go:generate go run github.com/kevinburke/go-bindata/go-bindata -pkg migrations -o ./migrations.bindata.go -ignore=\\*.go ./...

// Migrate calls db connections method to apply migration based on generated sql files
func Migrate(connection postgresql.Connection, version uint, log *logger.Logger) error {
	files, err := AssetDir("")
	if err != nil {
		return err
	}

	afn := func(name string) ([]byte, error) {
		data, err := Asset(name)
		if err != nil {
			return nil, err
		}
		return data, nil
	}

	migrationVersion, err := connection.Migrate(version, files, afn)
	if err != nil {
		return err
	}

	log.Info("migration applied to service",
		zap.Uint("version", migrationVersion))

	return nil
}
