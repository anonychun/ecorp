package entity

import (
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

type UserSession struct {
	Base

	UserId    uuid.UUID
	User      *User
	Token     string
	IpAddress string
	UserAgent string
}

func (as *UserSession) GenerateToken() {
	as.Token = ulid.Make().String()
}
