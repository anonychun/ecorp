package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuidv7()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Base) BeforeUpdate(tx *gorm.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}
