package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id             uuid.UUID `gorm:"type:uuid;default:uuidv7()"`
	Name           string
	EmailAddress   string
	PasswordDigest string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(hash)

	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
}
