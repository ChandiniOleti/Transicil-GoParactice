package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"authservice/generated"
	"authservice/services"
)

// AdminResponse is returned to clients without password hashes.
type AdminResponse struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetAdmins returns the list of all admins.
func GetAdmins(c *gin.Context) {
	admins, err := services.GetAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := make([]AdminResponse, 0, len(admins))
	for _, admin := range admins {
		response = append(response, AdminResponse{
			ID:    admin.ID,
			Name:  admin.Name,
			Email: admin.Email,
		})
	}

	c.JSON(http.StatusOK, response)
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
