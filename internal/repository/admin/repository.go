package admin

import (
	"context"

	"github.com/anonychun/bibit/internal/entity"
)

func (r *Repository) FindById(ctx context.Context, id string) (*entity.Admin, error) {
	admin := &entity.Admin{}
	err := r.sql.DB(ctx).First(admin, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *Repository) FindByEmailAddress(ctx context.Context, emailAddress string) (*entity.Admin, error) {
	admin := &entity.Admin{}
	err := r.sql.DB(ctx).First(admin, "email_address = ?", emailAddress).Error
	if err != nil {
		return nil, err
	}

	return admin, nil
}
