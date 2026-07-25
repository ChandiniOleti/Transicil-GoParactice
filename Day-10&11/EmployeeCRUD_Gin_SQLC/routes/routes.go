package routes

import (
	"employeecrudsqlc/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.POST("/students", handlers.CreateStudent)

	router.GET("/students", handlers.GetStudents)

}