package entity

import "golang.org/x/crypto/bcrypt"

type Admin struct {
	Base

	Name           string
	EmailAddress   string
	PasswordDigest string
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
