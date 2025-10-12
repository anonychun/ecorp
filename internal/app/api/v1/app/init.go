package app

import (
	"github.com/anonychun/benih/internal/app/api/v1/app/auth"
	"github.com/anonychun/benih/internal/bootstrap"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewUsecase)
	do.Provide(bootstrap.Injector, NewHandler)
}

type Usecase struct {
	Auth *auth.Usecase
}

func NewUsecase(i do.Injector) (*Usecase, error) {
	return &Usecase{
		Auth: do.MustInvoke[*auth.Usecase](i),
	}, nil
}

type Handler struct {
	Auth *auth.Handler
}

func NewHandler(i do.Injector) (*Handler, error) {
	return &Handler{
		Auth: do.MustInvoke[*auth.Handler](i),
	}, nil
}
