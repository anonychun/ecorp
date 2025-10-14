package v1

import (
	"github.com/anonychun/bibit/internal/app/api/v1/admin"
	"github.com/anonychun/bibit/internal/app/api/v1/app"
	"github.com/anonychun/bibit/internal/bootstrap"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewUsecase)
	do.Provide(bootstrap.Injector, NewHandler)
}

type Usecase struct {
	Admin *admin.Usecase
	App   *app.Usecase
}

func NewUsecase(i do.Injector) (*Usecase, error) {
	return &Usecase{
		Admin: do.MustInvoke[*admin.Usecase](i),
		App:   do.MustInvoke[*app.Usecase](i),
	}, nil
}

type Handler struct {
	Admin *admin.Handler
	App   *app.Handler
}

func NewHandler(i do.Injector) (*Handler, error) {
	return &Handler{
		Admin: do.MustInvoke[*admin.Handler](i),
		App:   do.MustInvoke[*app.Handler](i),
	}, nil
}
