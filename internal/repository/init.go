package repository

import (
	"context"

	"github.com/anonychun/bibit/internal/bootstrap"
	"github.com/anonychun/bibit/internal/current"
	"github.com/anonychun/bibit/internal/db"
	"github.com/anonychun/bibit/internal/repository/admin"
	"github.com/anonychun/bibit/internal/repository/admin_session"
	"github.com/anonychun/bibit/internal/repository/user"
	"github.com/anonychun/bibit/internal/repository/user_session"
	"github.com/samber/do/v2"
	"gorm.io/gorm"
)

func init() {
	do.Provide(bootstrap.Injector, NewRepository)
}

type Repository struct {
	Admin        *admin.Repository
	AdminSession *admin_session.Repository
	User         *user.Repository
	UserSession  *user_session.Repository
}

func NewRepository(i do.Injector) (*Repository, error) {
	return &Repository{
		Admin:        do.MustInvoke[*admin.Repository](i),
		AdminSession: do.MustInvoke[*admin_session.Repository](i),
		User:         do.MustInvoke[*user.Repository](i),
		UserSession:  do.MustInvoke[*user_session.Repository](i),
	}, nil
}

func Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	sql, err := do.Invoke[*db.Sql](bootstrap.Injector)
	if err != nil {
		return err
	}

	return sql.DB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = current.SetTx(ctx, tx)
		return fn(ctx)
	})
}
