package admin

import (
	"github.com/anonychun/ecorp/internal/app/api/v1/admin/admin"
	"github.com/anonychun/ecorp/internal/app/api/v1/admin/auth"
	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/samber/do"
)

func init() {
	do.ProvideNamed(bootstrap.Injector, UsecaseInjectorName, NewUsecase)
	do.ProvideNamed(bootstrap.Injector, HandlerInjectorName, NewHandler)
}

const (
	UsecaseInjectorName = "usecase.api.v1.admin"
	HandlerInjectorName = "handler.api.v1.admin"
)

type Usecase struct {
	Admin *admin.Usecase
	Auth  *auth.Usecase
}

func NewUsecase(i *do.Injector) (*Usecase, error) {
	return &Usecase{
		Admin: do.MustInvokeNamed[*admin.Usecase](i, admin.UsecaseInjectorName),
		Auth:  do.MustInvokeNamed[*auth.Usecase](i, auth.UsecaseInjectorName),
	}, nil
}

type Handler struct {
	Auth  *auth.Handler
	Admin *admin.Handler
}

func NewHandler(i *do.Injector) (*Handler, error) {
	return &Handler{
		Admin: do.MustInvokeNamed[*admin.Handler](i, admin.HandlerInjectorName),
		Auth:  do.MustInvokeNamed[*auth.Handler](i, auth.HandlerInjectorName),
	}, nil
}
