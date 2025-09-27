package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type UserSession struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuidv7()"`
	UserId    uuid.UUID
	User      *User
	Token     string
	IpAddress string
	UserAgent string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (as *UserSession) BeforeUpdate(tx *gorm.DB) error {
	as.UpdatedAt = time.Now()
	return nil
}

func (as *UserSession) GenerateToken() {
	as.Token = ulid.Make().String()
}
