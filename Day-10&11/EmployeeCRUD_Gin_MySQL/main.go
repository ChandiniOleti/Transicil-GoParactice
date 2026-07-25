package main

import (
	"employeecrud/config"
	"employeecrud/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}