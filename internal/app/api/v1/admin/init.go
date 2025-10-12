package admin

import (
	"github.com/anonychun/benih/internal/app/api/v1/admin/admin"
	"github.com/anonychun/benih/internal/app/api/v1/admin/auth"
	"github.com/anonychun/benih/internal/bootstrap"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewUsecase)
	do.Provide(bootstrap.Injector, NewHandler)
}

type Usecase struct {
	Admin *admin.Usecase
	Auth  *auth.Usecase
}

func NewUsecase(i do.Injector) (*Usecase, error) {
	return &Usecase{
		Admin: do.MustInvoke[*admin.Usecase](i),
		Auth:  do.MustInvoke[*auth.Usecase](i),
	}, nil
}

type Handler struct {
	Auth  *auth.Handler
	Admin *admin.Handler
}

func NewHandler(i do.Injector) (*Handler, error) {
	return &Handler{
		Admin: do.MustInvoke[*admin.Handler](i),
		Auth:  do.MustInvoke[*auth.Handler](i),
	}, nil
}
