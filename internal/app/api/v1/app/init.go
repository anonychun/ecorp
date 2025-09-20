package app

import (
	"github.com/anonychun/ecorp/internal/app/api/v1/app/auth"
	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/samber/do"
)

func init() {
	do.ProvideNamed(bootstrap.Injector, UsecaseInjectorName, NewUsecase)
	do.ProvideNamed(bootstrap.Injector, HandlerInjectorName, NewHandler)
}

const (
	UsecaseInjectorName = "usecase.api.v1.app"
	HandlerInjectorName = "handler.api.v1.app"
)

type Usecase struct {
	Auth *auth.Usecase
}

func NewUsecase(i *do.Injector) (*Usecase, error) {
	return &Usecase{
		Auth: do.MustInvokeNamed[*auth.Usecase](i, auth.UsecaseInjectorName),
	}, nil
}

type Handler struct {
	Auth *auth.Handler
}

func NewHandler(i *do.Injector) (*Handler, error) {
	return &Handler{
		Auth: do.MustInvokeNamed[*auth.Handler](i, auth.HandlerInjectorName),
	}, nil
}
