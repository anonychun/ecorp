package api

import (
	v1 "github.com/anonychun/bibit/internal/app/api/v1"
	"github.com/anonychun/bibit/internal/bootstrap"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewUsecase)
	do.Provide(bootstrap.Injector, NewHandler)
}

type Usecase struct {
	V1 *v1.Usecase
}

func NewUsecase(i do.Injector) (*Usecase, error) {
	return &Usecase{
		V1: do.MustInvoke[*v1.Usecase](i),
	}, nil
}

type Handler struct {
	V1 *v1.Handler
}

func NewHandler(i do.Injector) (*Handler, error) {
	return &Handler{
		V1: do.MustInvoke[*v1.Handler](i),
	}, nil
}
