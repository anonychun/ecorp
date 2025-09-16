package current

import (
	"context"

	"github.com/anonychun/ecorp/internal/entity"
	"gorm.io/gorm"
)

type key int

const (
	txKey key = iota
	adminKey
)

func Tx(ctx context.Context) *gorm.DB {
	tx, _ := ctx.Value(txKey).(*gorm.DB)
	return tx
}

func SetTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey, tx)
}

func Admin(ctx context.Context) *entity.Admin {
	admin, _ := ctx.Value(adminKey).(*entity.Admin)
	return admin
}

func SetAdmin(ctx context.Context, admin *entity.Admin) context.Context {
	return context.WithValue(ctx, adminKey, admin)
}
