package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"paylaterservice/services"
)


func LoginHandler(c *gin.Context) {

	var request services.LoginRequest


	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}


	token, err := services.Login(request)


	if err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})

		return
	}


	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}