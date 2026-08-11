package handlers

import (
	"github.com/gin-gonic/gin"

	"paybackservice/services"
)

// Payback processes a user's payback request.
// POST /payback
func Payback(c *gin.Context) {
	var request services.PaybackRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	role := c.GetString("role")
	tokenUserID := int32(c.GetInt("user_id"))

	// USER may only pay back their own account; ADMIN may pay back any user.
	if role != "ADMIN" && role != "USER" {
		c.JSON(403, gin.H{"error": "User access only"})
		return
	}
	if role == "USER" && tokenUserID != request.UserID {
		c.JSON(403, gin.H{"error": "You can access only your own account"})
		return
	}

	authHeader := c.GetHeader("Authorization")
	response, err := services.ProcessPayback(request, authHeader)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, response)
}
