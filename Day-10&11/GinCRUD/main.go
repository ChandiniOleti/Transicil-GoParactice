// package main

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// func main() {

// 	router := gin.Default()

// 	// Home Route
// 	router.GET("/", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "Hello Gin",
// 		})
// 	})

// 	// Hello Route
// 	router.GET("/hello", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "Hello World",
// 		})
// 	})

// 	// POST Request
// 	router.POST("/employee", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "Employee Created",
// 		})
// 	})

// 	// PUT Request
// 	router.PUT("/employee", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "Employee Updated",
// 		})
// 	})

// 	// DELETE Request
// 	router.DELETE("/employee", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "Employee Deleted",
// 		})
// 	})

// 	// Path Parameter
// 	router.GET("/employee/:id", func(c *gin.Context) {

// 		id := c.Param("id")

// 		c.JSON(200, gin.H{
// 			"id": id,
// 		})
// 	})

// 	// Multiple Path Parameters
// 	router.GET("/employee/:id/:department", func(c *gin.Context) {

// 		id := c.Param("id")
// 		department := c.Param("department")

// 		c.JSON(200, gin.H{
// 			"id":         id,
// 			"department": department,
// 		})
// 	})

// 	// String to Integer Conversion
// 	router.GET("/empint/:id", func(c *gin.Context) {

// 		id := c.Param("id")

// 		empID, err := strconv.Atoi(id)

// 		if err != nil {
// 			c.JSON(400, gin.H{
// 				"error": "Invalid Employee ID",
// 			})
// 			return
// 		}

// 		fmt.Println("Employee ID:", empID)

// 		c.JSON(200, gin.H{
// 			"employee_id": empID,
// 		})
// 	})

// 	// Query Parameters
// 	router.GET("/search", func(c *gin.Context) {

// 		name := c.Query("name")
// 		city := c.Query("city")

// 		c.JSON(200, gin.H{
// 			"name": name,
// 			"city": city,
// 		})
// 	})

// 	// Query Parameter with Default Value
// 	router.GET("/employees", func(c *gin.Context) {

// 	city := c.Query("city")
// 	dept := c.Query("department")

// 	c.JSON(200, gin.H{
// 		"city": city,
// 		"department": dept,
// 	})

// })

// 	router.Run(":8080")
// }



package main

import "github.com/gin-gonic/gin"

type Employee struct {
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

func main() {

	router := gin.Default()

	router.POST("/employee", func(c *gin.Context) {

		var employee Employee

		err := c.BindJSON(&employee)

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid JSON",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Employee Created",
			"data":    employee,
		})

	})

	router.Run(":8080")
}