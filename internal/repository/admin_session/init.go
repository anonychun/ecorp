package admin_session

import (
	"github.com/anonychun/benih/internal/bootstrap"
	"github.com/anonychun/benih/internal/db"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewRepository)
}

type Repository struct {
	sql *db.Sql
}

func NewRepository(i do.Injector) (*Repository, error) {
	return &Repository{
		sql: do.MustInvoke[*db.Sql](i),
	}, nil
}
