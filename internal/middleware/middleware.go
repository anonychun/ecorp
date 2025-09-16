package middleware

import (
	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/anonychun/ecorp/internal/middleware/auth"
	"github.com/samber/do"
)

func init() {
	do.Provide(bootstrap.Injector, NewMiddleware)
}

type Middleware struct {
	Auth *auth.Middleware
}

func NewMiddleware(i *do.Injector) (*Middleware, error) {
	return &Middleware{
		Auth: do.MustInvokeNamed[*auth.Middleware](i, auth.MiddlewareInjectorName),
	}, nil
}
