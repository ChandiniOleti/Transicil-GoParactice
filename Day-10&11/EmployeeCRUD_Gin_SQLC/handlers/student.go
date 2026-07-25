package handlers

import (
	"context"
	"net/http"

	"employeecrudsqlc/config"
	"employeecrudsqlc/generated"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {

	var req generated.CreateStudentParams

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := config.Queries.CreateStudent(context.Background(), req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Student created successfully",
	})
}


func GetStudents(c *gin.Context) {

	students, err := config.Queries.GetStudents(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, students)
}