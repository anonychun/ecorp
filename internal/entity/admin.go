package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	Id             uuid.UUID
	Name           string
	EmailAddress   string
	PasswordDigest string
	DeletedAt      *time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (a *Admin) BeforeCreate(tx *gorm.DB) error {
	a.Id = uuid.Must(uuid.NewV7())
	return nil
}

func (a *Admin) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.PasswordDigest = string(hash)

	return nil
}

func (a *Admin) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(a.PasswordDigest), []byte(password))
}
