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

func (r *Repository) DeleteById(ctx context.Context, id string) error {
	return r.sql.DB(ctx).Delete(&entity.AdminSession{}, "id = ?", id).Error
}

func (r *Repository) DeleteAllByAdminId(ctx context.Context, adminId string) error {
	return r.sql.DB(ctx).Delete(&entity.AdminSession{}, "admin_id = ?", adminId).Error
}
