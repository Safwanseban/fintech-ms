package controllers

import (
	pkg "fintechGo/internal/pkg"
	"fintechGo/internal/types"
	services "fintechGo/internal/usecases/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Userhandler struct {
	userUsercase services.UserInterface
}

func NewHandler(ctx *gin.Engine, user services.UserInterface) {
	handler := &Userhandler{
		userUsercase: user,
	}
	ctx.POST("/", handler.CreateUser)

}

func (u *Userhandler) CreateUser(ctx *gin.Context) {
	var users *types.AuthUser
	if err := ctx.ShouldBindJSON(&users); err != nil {
		pkg.ErrorResponse(ctx, http.StatusNotFound, err)
		return
	}
	jwt, err := u.userUsercase.CreateUser(users)
	if err != nil {
		pkg.ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(200, gin.H{
		"message":  "success",
		"token":    jwt,
	})

}
