package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"userservice/generated"
	"userservice/services"
)

// CreateUser creates a new user.
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

// GetUsers returns all users.
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

// GetUserByID returns a single user based on ID.
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

// UpdateUser updates an existing user's information.
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

// DeleteUser deletes a user using the given ID.
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

// UpdateCurrentDueRequest is the body for internal due updates.
type UpdateCurrentDueRequest struct {
	CurrentDue string `json:"current_due"`
}

// UpdateCurrentDue updates a user's current_due (internal only).
// PATCH /internal/users/:id/due
func UpdateCurrentDue(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid User ID",
		})
		return
	}

	var request UpdateCurrentDueRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = services.UpdateCurrentDue(int32(id), request.CurrentDue)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Current due updated successfully",
	})
}

// GetUsersInternal returns all users for reporting (internal only).
// GET /internal/users
func GetUsersInternal(c *gin.Context) {
	users, err := services.GetUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}

// GetUserByIDInternal returns one user for reporting (internal only).
// GET /internal/users/:id
func GetUserByIDInternal(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid User ID"})
		return
	}

	user, err := services.GetUserByID(int32(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "User Not Found"})
		return
	}
	c.JSON(200, user)
}
