package main

import (
	"fintechGo/configs"
	"fintechGo/internal/controllers"
	"fintechGo/internal/repo"
	"fintechGo/internal/usecases"
	"log"

	"github.com/gin-gonic/gin"
)


func main() {

	conf := configs.NewConfig()
	db := configs.ConnectToDB(conf)
	engine := gin.Default()
	repo := repo.NewUser(db)
	usecase := usecases.NewUserUseCase(repo)
	controllers.NewHandler(engine, usecase)

	if err := engine.Run(conf.String("port")); err != nil {
		log.Fatalf("err %v", err)
	}
}
