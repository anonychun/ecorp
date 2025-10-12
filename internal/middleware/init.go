package middleware

import (
	"github.com/anonychun/benih/internal/bootstrap"
	"github.com/anonychun/benih/internal/middleware/auth"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewMiddleware)
}

type Middleware struct {
	Auth *auth.Middleware
}

func NewMiddleware(i do.Injector) (*Middleware, error) {
	return &Middleware{
		Auth: do.MustInvoke[*auth.Middleware](i),
	}, nil
}
