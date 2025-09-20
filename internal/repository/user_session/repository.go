package user_session

import (
	"context"

	"github.com/anonychun/ecorp/internal/entity"
)

func (r *Repository) FindByToken(ctx context.Context, token string) (*entity.UserSession, error) {
	userSession := &entity.UserSession{}
	err := r.sql.DB(ctx).First(userSession, "token = ?", token).Error
	if err != nil {
		return nil, err
	}

	return userSession, nil
}

func (r *Repository) Create(ctx context.Context, userSession *entity.UserSession) error {
	return r.sql.DB(ctx).Create(userSession).Error
}

func (r *Repository) DeleteById(ctx context.Context, id string) error {
	return r.sql.DB(ctx).Delete(&entity.UserSession{}, "id = ?", id).Error
}
