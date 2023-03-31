package controllers

import (
	"errors"
	pkg "fintechGo/internal/pkg"
	"fintechGo/internal/types"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *Userhandler) CreateUser(ctx *gin.Context) {
	var users *types.AuthUser
	if err := ctx.ShouldBindJSON(&users); err != nil {
		pkg.ErrorResponse(ctx, http.StatusNotFound, errors.New("bad payload"))
		return
	}
	if err := u.userUsercase.CreateUser(users); err != nil {
		pkg.ErrorResponse(ctx, http.StatusInternalServerError, errors.New("error encountered"))
		return
	}
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}
