package controllers

import (
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
	ctx.GET("/", handler.GetData)

}

func (u *Userhandler) GetData(ctx *gin.Context) {
	var users *types.AuthUser
	err := u.userUsercase.CreateUser(users)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	ctx.JSON(200, gin.H{
		"message": "hai",
	})

}
