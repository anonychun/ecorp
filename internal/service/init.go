package service

import (
	"github.com/anonychun/bibit/internal/bootstrap"
	"github.com/samber/do/v2"
)

func init() {
	do.Provide(bootstrap.Injector, NewService)
}

type Service struct {
}

func NewService(i do.Injector) (*Service, error) {
	return &Service{}, nil
}
