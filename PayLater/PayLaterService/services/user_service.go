package services

import (
	"context"

	"paylaterservice/config"
	"paylaterservice/generated"
	"golang.org/x/crypto/bcrypt"
)

// UserResponse defines the JSON response returned to the client.
// Password is intentionally excluded for security reasons.
type UserResponse struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	CreditLimit string `json:"credit_limit"`
	CurrentDue  string `json:"current_due"`
}

// CreateUser creates a new user in the database.
//
// Example:
// Name     : Chandini
// Email    : chandini@gmail.com
// Password : 123456
// Role     : USER
//
// Before storing the user, the plain-text password is
// converted into a secure hashed password using bcrypt.
func CreateUser(user generated.CreateUserParams) error {

	// Encrypt the password before saving it to the database.
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	// Replace the original password with the hashed password.
	user.Password = string(hashedPassword)

	// Insert the user into the database.
	_, err = config.Queries.CreateUser(context.Background(), user)

	return err
}

// GetUsers retrieves all users from the database.
//
// Password is not returned in the response for security reasons.
func GetUsers() ([]UserResponse, error) {

	users, err := config.Queries.GetUsers(context.Background())

	if err != nil {
		return nil, err
	}

	var response []UserResponse

	// Convert database model into API response.
	for _, user := range users {

		response = append(response, UserResponse{
	ID:          user.ID,
	Name:        user.Name,
	Email:       user.Email,
	CreditLimit: user.CreditLimit,
	CurrentDue:  user.CurrentDue,
})
	}

	return response, nil
}

// GetUserByID retrieves a specific user using the user ID.
//
// Example:
// GET /users/1
func GetUserByID(id int32) (UserResponse, error) {

	user, err := config.Queries.GetUserByID(context.Background(), id)

	if err != nil {
		return UserResponse{}, err
	}

	// Prepare response without exposing the password.
	response := UserResponse{
	ID:          user.ID,
	Name:        user.Name,
	Email:       user.Email,
	CreditLimit: user.CreditLimit,
	CurrentDue:  user.CurrentDue,
}

	return response, nil
}

// UpdateUser updates an existing user's information.
//
// If the password is changed, it is hashed again before
// storing it in the database.
func UpdateUser(user generated.UpdateUserParams) error {

	// Encrypt the updated password.
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	// Store the encrypted password.
	user.Password = string(hashedPassword)

	// Update the user details in the database.
	return config.Queries.UpdateUser(context.Background(), user)
}

// DeleteUser removes a user from the database using the user ID.
//
// Example:
// DELETE /users/5
func DeleteUser(id int32) error {

	err := config.Queries.DeleteUser(context.Background(), id)

	if err != nil {
		return err
	}

	return nil
}