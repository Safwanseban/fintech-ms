package types

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthUser struct {
	Name     string
	Email    string
	Password string
}

func (user *AuthUser) HashPassword(password string) (string, error) {
	if password == "" {

		return "", errors.New("no password provided")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (user *AuthUser) CheckPassword(password, actual string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(actual))

	return err == nil

}
