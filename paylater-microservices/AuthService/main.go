package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"authservice/config"
	"authservice/routes"
	"authservice/utils"

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
	routes.SetupRoutes(router, rate.Limit(2), 5)

	if err := server.Run(config.ListenAddr(), router, config.DB); err != nil {
		log.Fatal(err)
	}
}
