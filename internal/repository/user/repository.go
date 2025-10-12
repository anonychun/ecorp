package user

import (
	"context"

	"github.com/anonychun/benih/internal/entity"
)

func (r *Repository) FindById(ctx context.Context, id string) (*entity.User, error) {
	user := &entity.User{}
	err := r.sql.DB(ctx).First(user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) FindByEmailAddress(ctx context.Context, emailAddress string) (*entity.User, error) {
	user := &entity.User{}
	err := r.sql.DB(ctx).First(user, "email_address = ?", emailAddress).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) Create(ctx context.Context, user *entity.User) error {
	return r.sql.DB(ctx).Create(user).Error
}

func (r *Repository) ExistsByEmailAddress(ctx context.Context, emailAddress string) (bool, error) {
	var exists bool
	err := r.sql.DB(ctx).Raw("SELECT 1 FROM users WHERE email_address = ?", emailAddress).Scan(&exists).Error
	return exists, err
}
