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

func (user *AuthUser) CheckPassword(actual, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(actual), []byte(password))
	return err == nil

}
