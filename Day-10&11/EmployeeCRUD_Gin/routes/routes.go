package routes

import (
	"employeecrudgin/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/employees", handlers.GetEmployees)

	router.GET("/employees/:id", handlers.GetEmployeeByID)

	router.POST("/employees", handlers.CreateEmployee)

	router.PUT("/employees/:id", handlers.UpdateEmployee)

	router.DELETE("/employees/:id", handlers.DeleteEmployee)

}