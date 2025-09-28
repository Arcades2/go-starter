package userservice

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher struct{}

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{}
}

func (p *PasswordHasher) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (p *PasswordHasher) VerifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
