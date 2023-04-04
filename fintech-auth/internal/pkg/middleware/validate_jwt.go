package middleware

import (
	"errors"
	"fintechGo/configs"
	"fintechGo/internal/pkg"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	logger = configs.Getlogger()
)

func ValidateJwt() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenString := GetJWT(ctx)
		fmt.Println("pooyy")
		token, err := jwt.ParseWithClaims(tokenString, &UserJWT{}, func(t *jwt.Token) (interface{}, error) {

			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {

				return nil, fmt.Errorf("unexpected sigining method:%v", t.Header["alg"])
			}

			return []byte(secretKey), nil

		}, jwt.WithJSONNumber())
		logger.Error().Err(err).Send()
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
	fmt.Println(data)
	return data
	//return strings.Split(data, ".")

}
