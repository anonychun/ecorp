package user_session

import (
	"context"

	"github.com/anonychun/bibit/internal/entity"
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

func (r *Repository) DeleteByToken(ctx context.Context, token string) error {
	return r.sql.DB(ctx).Delete(&entity.UserSession{}, "token = ?", token).Error
}
