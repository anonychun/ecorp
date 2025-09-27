package entity

import "golang.org/x/crypto/bcrypt"

type User struct {
	Base

	Name           string
	EmailAddress   string
	PasswordDigest string
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
