package controllers

import (
	"errors"
	"fintechGo/internal/pkg"
	"fintechGo/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *Userhandler) ValidateUser(ctx *gin.Context) {
	var user *types.AuthUser
	if err := ctx.ShouldBindJSON(&user); err != nil {
		pkg.ErrorResponse(ctx, http.StatusBadRequest, errors.New("bad payload data"))
		return
	}
	jwt, err := u.userUsercase.ValidateUser(user)
	if err != nil {
		pkg.ErrorResponse(ctx, http.StatusInternalServerError, errors.New("error encountered"))
		return
	}
	ctx.Header("UserJWT", jwt["userJWT"])
	ctx.JSON(http.StatusOK, gin.H{
		"message": "succesfull login",
		"jwt":     jwt,
	})
}
