package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"paylaterservice/generated"
	"paylaterservice/services"
)

func GetAdmins(c *gin.Context) {

	admins, err := services.GetAdmins()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, admins)
}



func CreateAdmin(c *gin.Context) {

	var admin generated.CreateAdminParams

	if err := c.ShouldBindJSON(&admin); err != nil {

		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := services.CreateAdmin(admin)

	if err != nil {

		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "Admin Created Successfully",
	})
}