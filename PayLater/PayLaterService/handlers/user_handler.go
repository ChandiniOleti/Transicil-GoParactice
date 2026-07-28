package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"paylaterservice/generated"
	"paylaterservice/services"
)

// POST /users
func CreateUser(c *gin.Context) {

	var user generated.CreateUserParams

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := services.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "User Created Successfully",
	})
}

// GET /users
func GetUsers(c *gin.Context) {

	users, err := services.GetUsers()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

// GET /users/:id
func GetUserByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid User ID",
		})
		return
	}

	user, err := services.GetUserByID(int32(id))
	if err != nil {
		c.JSON(404, gin.H{
			"error": "User Not Found",
		})
		return
	}

	c.JSON(200, user)
}

// PUT /users/:id
func UpdateUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid User ID",
		})
		return
	}

	var user generated.UpdateUserParams

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.ID = int32(id)

	err = services.UpdateUser(user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User Updated Successfully",
	})
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid User ID",
		})
		return
	}

	err = services.DeleteUser(int32(id))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User Deleted Successfully",
	})
}