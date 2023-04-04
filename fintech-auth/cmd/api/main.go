package main

import (
	"fintechGo/internal/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	usecase, conf := controllers.Initialize()
	controllers.NewHandler(engine, usecase)

	if err := engine.Run(conf.String("port")); err != nil {
		log.Fatalf("err %v", err)
	}
}
