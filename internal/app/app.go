package app

import (
	"github.com/anonychun/ecorp/internal/app/api"
	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/samber/do"
)

func init() {
	do.Provide(bootstrap.Injector, NewUsecase)
	do.Provide(bootstrap.Injector, NewHandler)
}

type Usecase struct {
	Api *api.Usecase
}

func NewUsecase(i *do.Injector) (*Usecase, error) {
	return &Usecase{
		Api: do.MustInvokeNamed[*api.Usecase](i, api.UsecaseInjectorName),
	}, nil
}

type Handler struct {
	Api *api.Handler
}

func NewHandler(i *do.Injector) (*Handler, error) {
	return &Handler{
		Api: do.MustInvokeNamed[*api.Handler](i, api.HandlerInjectorName),
	}, nil
}
