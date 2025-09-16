package db

import (
	"context"

	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/anonychun/ecorp/migrations"
	"github.com/pressly/goose/v3"
	"github.com/samber/do"
)

func init() {
	do.Provide(bootstrap.Injector, NewMigrator)
}

type Migrator struct {
	provider *goose.Provider
}

func NewMigrator(i *do.Injector) (*Migrator, error) {
	provider, err := goose.NewProvider(
		"postgres",
		do.MustInvoke[*Sql](i).sqlDB,
		migrations.MigrationsFs,
		goose.WithVerbose(true),
	)
	if err != nil {
		return nil, err
	}

	return &Migrator{
		provider: provider,
	}, nil
}

func (m *Migrator) Migrate(ctx context.Context) error {
	_, err := m.provider.Up(ctx)
	return err
}

func (m *Migrator) Rollback(ctx context.Context) error {
	_, err := m.provider.Down(ctx)
	return err
}
