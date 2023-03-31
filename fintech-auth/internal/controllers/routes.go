package controllers

import (
	services "fintechGo/internal/usecases/interfaces"

	"github.com/gin-gonic/gin"
)

type Userhandler struct {
	userUsercase services.UserInterface
}

func NewHandler(ctx *gin.Engine, user services.UserInterface) {
	handler := &Userhandler{
		userUsercase: user,
	}
	ctx.POST("/register", handler.CreateUser)
	ctx.POST("/login", handler.ValidateUser)

}
