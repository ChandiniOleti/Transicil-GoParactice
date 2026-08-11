package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"authservice/generated"
	"authservice/services"
)

// GetAdmins returns the list of all admins.
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

// CreateAdmin creates a new admin.
func CreateAdmin(c *gin.Context) {
	var admin generated.CreateAdminParams

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := services.CreateAdmin(admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Admin Created Successfully",
	})
}
