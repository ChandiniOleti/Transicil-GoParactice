package routes

import (
	"employeecrud/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {


	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Employee CRUD API is running",
		})
	})


	// POST API
	router.POST("/employees", handlers.CreateEmployee)
	router.GET("/employees", handlers.GetEmployees)
	router.GET("/employees/:id", handlers.GetEmployeeByID)
	router.PUT("/employees/:id", handlers.UpdateEmployee)
	router.DELETE("/employees/:id", handlers.DeleteEmployee)
}
