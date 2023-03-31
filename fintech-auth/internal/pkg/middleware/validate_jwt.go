package middleware

import (
	"errors"
	"fintechGo/internal/pkg"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ValidateJwt() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenString := GetJWT(ctx)

		token, err := jwt.ParseWithClaims(tokenString, UserJWT{}, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected sigining method:%v", t.Header["alg"])
			}

			return []byte(secretKey), nil
		}, nil)

		if !token.Valid {
			switch {

			case errors.Is(err, jwt.ErrTokenExpired):
				pkg.ErrorResponse(ctx, http.StatusUnauthorized, errors.New("token expired"))
				ctx.Abort()
				return
			case errors.Is(err, jwt.ErrTokenMalformed):
				pkg.ErrorResponse(ctx, http.StatusUnauthorized, errors.New("malformed token"))
				ctx.Abort()
				return
			}
			pkg.ErrorResponse(ctx, http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}

		ctx.Set("JwtData", token.Claims)
		ctx.Next()
	}
}
func GetJWT(ctx *gin.Context) string {

	data := ctx.GetHeader("UserJWT")
	return data
	//return strings.Split(data, ".")

}
