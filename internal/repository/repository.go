package repository

import (
	"context"

	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/anonychun/ecorp/internal/current"
	"github.com/anonychun/ecorp/internal/db"
	"github.com/anonychun/ecorp/internal/repository/admin"
	"github.com/anonychun/ecorp/internal/repository/admin_session"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func init() {
	do.Provide(bootstrap.Injector, NewRepository)
}

type Repository struct {
	sql *db.Sql

	Admin        *admin.Repository
	AdminSession *admin_session.Repository
}

func NewRepository(i *do.Injector) (*Repository, error) {
	return &Repository{
		sql: do.MustInvoke[*db.Sql](i),

		Admin:        do.MustInvokeNamed[*admin.Repository](i, admin.RepositoryInjectorName),
		AdminSession: do.MustInvokeNamed[*admin_session.Repository](i, admin_session.RepositoryInjectorName),
	}, nil
}

func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.sql.DB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = current.SetTx(ctx, tx)
		return fn(ctx)
	})
}
