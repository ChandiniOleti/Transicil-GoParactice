package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"transactionservice/config"
	"transactionservice/routes"
	"transactionservice/utils"

	"paylater.dev/shared/logger"
	"paylater.dev/shared/requestid"
	"paylater.dev/shared/server"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	logger.Init(config.ServiceName)
	utils.InitJWT(config.JWTSecret)

	if err := config.ConnectDB(); err != nil {
		log.Fatal("Database Connection Failed:", err)
	}
	config.InitSQLC()

	router := gin.New()
	router.Use(gin.Recovery(), requestid.Middleware(), logger.GinMiddleware())
	routes.SetupRoutes(router)

	if err := server.Run(config.ListenAddr(), router, config.DB); err != nil {
		log.Fatal(err)
	}
}
