package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type UserSession struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	User      *User
	Token     string
	IpAddress string
	UserAgent string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (as *UserSession) BeforeCreate(tx *gorm.DB) error {
	as.Id = uuid.Must(uuid.NewV7())
	return nil
}

func (as *UserSession) BeforeUpdate(tx *gorm.DB) error {
	as.UpdatedAt = time.Now()
	return nil
}

func (as *UserSession) GenerateToken() {
	as.Token = ulid.Make().String()
}
