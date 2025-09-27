package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type AdminSession struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuidv7()"`
	AdminId   uuid.UUID
	Admin     *Admin
	Token     string
	IpAddress string
	UserAgent string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (as *AdminSession) BeforeUpdate(tx *gorm.DB) error {
	as.UpdatedAt = time.Now()
	return nil
}

func (as *AdminSession) GenerateToken() {
	as.Token = ulid.Make().String()
}
