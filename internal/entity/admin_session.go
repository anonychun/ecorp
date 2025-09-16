package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type AdminSession struct {
	Id        uuid.UUID
	AdminId   uuid.UUID
	Token     string
	IpAddress string
	UserAgent string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (as *AdminSession) BeforeCreate(tx *gorm.DB) error {
	as.Id = uuid.Must(uuid.NewV7())
	return nil
}

func (as *AdminSession) GenerateToken() {
	as.Token = ulid.Make().String()
}
