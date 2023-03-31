package pkg

import (
	"fintechGo/configs"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserJWT struct {
	Name  string
	Email string
	jwt.RegisteredClaims
}

var (
	conf      = configs.NewConfig()
	secretKey = conf.Bytes("jwt.secret")
)

func CreateJWT(email string) (map[string]string, error) {
	claims := UserJWT{

		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"userJWT": signedToken,
	}, nil

}
