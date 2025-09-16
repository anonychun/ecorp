package db

import (
	"context"

	"github.com/anonychun/ecorp/internal/bootstrap"
	"github.com/anonychun/ecorp/internal/entity"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func init() {
	do.Provide(bootstrap.Injector, NewSeeder)
}

type Seeder struct {
	sql *Sql
}

func NewSeeder(i *do.Injector) (*Seeder, error) {
	return &Seeder{
		sql: do.MustInvoke[*Sql](i),
	}, nil
}

func (s *Seeder) Seed(ctx context.Context) error {
	defaultAdmin := &entity.Admin{
		Name:         "Achmad Chun Chun",
		EmailAddress: "anonychun@gmail.com",
	}

	defaultAdminPassword := "didbnyaada"
	err := defaultAdmin.HashPassword(defaultAdminPassword)
	if err != nil {
		return err
	}

	err = s.sql.DB(ctx).First(defaultAdmin, "email_address = ?", defaultAdmin.EmailAddress).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	defaultAdmin.DeletedAt = nil
	err = s.sql.DB(ctx).Save(defaultAdmin).Error
	if err != nil {
		return err
	}

	return nil
}
