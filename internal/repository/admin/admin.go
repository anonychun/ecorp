package admin

import (
	"context"

	"github.com/anonychun/ecorp/internal/entity"
)

func (r *Repository) FindAll(ctx context.Context) ([]*entity.Admin, error) {
	admins := make([]*entity.Admin, 0)
	err := r.sql.DB(ctx).Find(&admins).Error
	if err != nil {
		return nil, err
	}

	return admins, nil
}

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

func (r *Repository) Create(ctx context.Context, admin *entity.Admin) error {
	return r.sql.DB(ctx).Create(admin).Error
}

func (r *Repository) Update(ctx context.Context, admin *entity.Admin) error {
	return r.sql.DB(ctx).Save(admin).Error
}

func (r *Repository) DeleteById(ctx context.Context, id string) error {
	return r.sql.DB(ctx).Delete(&entity.Admin{}, "id = ?", id).Error
}

func (r *Repository) ExistsById(ctx context.Context, id string) (bool, error) {
	var exists bool
	err := r.sql.DB(ctx).Raw("SELECT 1 FROM admins WHERE id = ?", id).Scan(&exists).Error
	return exists, err
}
