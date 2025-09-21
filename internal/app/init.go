package app

import (
	"github.com/anonychun/ecorp/internal/app/api"
	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewUsecase)
	do.Provide(bootstrap.Injector, NewHandler)
}

type Usecase struct {
	Api *api.Usecase
}

func NewUsecase(i do.Injector) (*Usecase, error) {
	return &Usecase{
		Api: do.MustInvoke[*api.Usecase](i),
	}, nil
}

type Handler struct {
	Api *api.Handler
}

func NewHandler(i do.Injector) (*Handler, error) {
	return &Handler{
		Api: do.MustInvoke[*api.Handler](i),
	}, nil
}
