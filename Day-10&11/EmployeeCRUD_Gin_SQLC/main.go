package main

import (
	"employeecrudsqlc/config"
	"employeecrudsqlc/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()
	config.InitQueries()

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}