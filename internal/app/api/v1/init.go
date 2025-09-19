package v1

import (
	"github.com/anonychun/ecorp/internal/app/api/v1/admin"
	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/samber/do"
)

func init() {
	do.ProvideNamed(bootstrap.Injector, UsecaseInjectorName, NewUsecase)
	do.ProvideNamed(bootstrap.Injector, HandlerInjectorName, NewHandler)
}

const (
	UsecaseInjectorName = "usecase.api.v1"
	HandlerInjectorName = "handler.api.v1"
)

type Usecase struct {
	Admin *admin.Usecase
}

func NewUsecase(i *do.Injector) (*Usecase, error) {
	return &Usecase{
		Admin: do.MustInvokeNamed[*admin.Usecase](i, admin.UsecaseInjectorName),
	}, nil
}

type Handler struct {
	Admin *admin.Handler
}

func NewHandler(i *do.Injector) (*Handler, error) {
	return &Handler{
		Admin: do.MustInvokeNamed[*admin.Handler](i, admin.HandlerInjectorName),
	}, nil
}
