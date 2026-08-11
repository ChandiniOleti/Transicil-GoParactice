package services

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"userservice/config"
	"userservice/generated"
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
func CreateUser(user generated.CreateUserParams) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	_, err = config.Queries.CreateUser(context.Background(), user)

	return err
}

// GetUsers retrieves all users from the database.
func GetUsers() ([]UserResponse, error) {
	users, err := config.Queries.GetUsers(context.Background())
	if err != nil {
		return nil, err
	}

	var response []UserResponse

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
func GetUserByID(id int32) (UserResponse, error) {
	user, err := config.Queries.GetUserByID(context.Background(), id)
	if err != nil {
		return UserResponse{}, err
	}

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
func UpdateUser(user generated.UpdateUserParams) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return config.Queries.UpdateUser(context.Background(), user)
}

// DeleteUser removes a user from the database using the user ID.
func DeleteUser(id int32) error {
	return config.Queries.DeleteUser(context.Background(), id)
}

// UpdateCurrentDue updates only the user's current_due field.
func UpdateCurrentDue(id int32, currentDue string) error {
	return config.Queries.UpdateCurrentDue(context.Background(), generated.UpdateCurrentDueParams{
		ID:         id,
		CurrentDue: currentDue,
	})
}
