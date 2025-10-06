package auth

import (
	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/anonychun/ecorp/internal/repository"
	"github.com/anonychun/ecorp/internal/validator"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewUsecase)
	do.Provide(bootstrap.Injector, NewHandler)
}

type Usecase struct {
	validator  *validator.Validator
	repository *repository.Repository
}

func NewUsecase(i do.Injector) (*Usecase, error) {
	return &Usecase{
		validator:  do.MustInvoke[*validator.Validator](i),
		repository: do.MustInvoke[*repository.Repository](i),
	}, nil
}

type Handler struct {
	usecase *Usecase
}

func NewHandler(i do.Injector) (*Handler, error) {
	return &Handler{
		usecase: do.MustInvoke[*Usecase](i),
	}, nil
}
