package controllers

import (
	"fintechGo/configs"
	"fintechGo/internal/pkg/middleware"
	"fintechGo/internal/repo"
	"fintechGo/internal/usecases"
	"fintechGo/internal/usecases/interfaces"
	services "fintechGo/internal/usecases/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf"
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
	ctx.GET("/", middleware.ValidateJwt(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ping": "hai",
		})
	})

}

func Initialize() (interfaces.UserInterface, *koanf.Koanf) {
	logger := configs.Getlogger()
	conf := configs.NewConfig()
	db := configs.ConnectToDB(conf)

	repo := repo.NewUser(db)
	usecase := usecases.NewUserUseCase(repo, logger)
	return usecase, conf
}
