package auth

import (
	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/anonychun/ecorp/internal/repository"
	"github.com/samber/do"
)

func init() {
	do.ProvideNamed(bootstrap.Injector, MiddlewareInjectorName, NewMiddleware)
}

const MiddlewareInjectorName = "middleware.auth"

type Middleware struct {
	repository *repository.Repository
}

func NewMiddleware(i *do.Injector) (*Middleware, error) {
	return &Middleware{
		repository: do.MustInvoke[*repository.Repository](i),
	}, nil
}
