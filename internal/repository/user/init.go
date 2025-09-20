package user

import (
	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/anonychun/ecorp/internal/db"
	"github.com/samber/do"
)

func init() {
	do.ProvideNamed(bootstrap.Injector, RepositoryInjectorName, NewRepository)
}

const RepositoryInjectorName = "repository.user"

type Repository struct {
	sql *db.Sql
}

func NewRepository(i *do.Injector) (*Repository, error) {
	return &Repository{
		sql: do.MustInvoke[*db.Sql](i),
	}, nil
}
