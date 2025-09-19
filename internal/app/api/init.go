package api

import (
	v1 "github.com/anonychun/ecorp/internal/app/api/v1"
	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/samber/do"
)

func init() {
	do.ProvideNamed(bootstrap.Injector, UsecaseInjectorName, NewUsecase)
	do.ProvideNamed(bootstrap.Injector, HandlerInjectorName, NewHandler)
}

const (
	UsecaseInjectorName = "usecase.api"
	HandlerInjectorName = "handler.api"
)

type Usecase struct {
	V1 *v1.Usecase
}

func NewUsecase(i *do.Injector) (*Usecase, error) {
	return &Usecase{
		V1: do.MustInvokeNamed[*v1.Usecase](i, v1.UsecaseInjectorName),
	}, nil
}

type Handler struct {
	V1 *v1.Handler
}

func NewHandler(i *do.Injector) (*Handler, error) {
	return &Handler{
		V1: do.MustInvokeNamed[*v1.Handler](i, v1.HandlerInjectorName),
	}, nil
}
