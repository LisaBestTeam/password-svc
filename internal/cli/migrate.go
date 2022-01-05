package cli

import (
	"github.com/lisabestteam/password-svc/internal/assets"
	"github.com/lisabestteam/password-svc/internal/config"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
)

var migrations = &migrate.PackrMigrationSource{
	Box: assets.Migrations,
}

func MigrateUp(cfg config.Config) error {
	applied, err := migrate.Exec(cfg.RawDB(), "postgres", migrations, migrate.Up)
	if err != nil {
		return errors.Wrap(err, "failed to apply migrations")
	}
	cfg.Log().WithField("applied", applied).Info("migrations applied")
	return nil
}

func MigrateDown(cfg config.Config) error {
	applied, err := migrate.Exec(cfg.RawDB(), "postgres", migrations, migrate.Down)
	if err != nil {
		return errors.Wrap(err, "failed to apply migrations")
	}
	cfg.Log().WithField("applied", applied).Info("migrations applied")
	return nil
}
