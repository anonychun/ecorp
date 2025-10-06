package admin_session

import (
	"context"

	"github.com/anonychun/ecorp/internal/entity"
)

func (r *Repository) FindByToken(ctx context.Context, token string) (*entity.AdminSession, error) {
	adminSession := &entity.AdminSession{}
	err := r.sql.DB(ctx).First(adminSession, "token = ?", token).Error
	if err != nil {
		return nil, err
	}

	return adminSession, nil
}

func (r *Repository) Create(ctx context.Context, adminSession *entity.AdminSession) error {
	return r.sql.DB(ctx).Create(adminSession).Error
}

func (r *Repository) DeleteByToken(ctx context.Context, token string) error {
	return r.sql.DB(ctx).Delete(&entity.AdminSession{}, "token = ?", token).Error
}
