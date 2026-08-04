package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"apigateway/config"
	"apigateway/routes"
	"apigateway/utils"

	"paylater.dev/shared/logger"
	"paylater.dev/shared/requestid"
	"paylater.dev/shared/server"
)

func main() {
	logger.Init("APIGateway")
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	logger.Init(config.ServiceName)
	utils.InitJWT(config.JWTSecret)

	router := gin.New()
	router.Use(gin.Recovery(), requestid.Middleware(), logger.GinMiddleware())
	routes.SetupRoutes(router)

	if err := server.Run(config.ListenAddr(), router, nil); err != nil {
		log.Fatal(err)
	}
}
