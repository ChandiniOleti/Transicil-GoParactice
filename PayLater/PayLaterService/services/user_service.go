package services

import (
	"context"

	"paylaterservice/config"
	"paylaterservice/generated"
)

type UserResponse struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	CreditLimit string `json:"credit_limit"`
	CurrentDue  string `json:"current_due"`
}

func CreateUser(user generated.CreateUserParams) error {

	_, err := config.Queries.CreateUser(context.Background(), user)

	if err != nil {
		return err
	}

	return nil
}


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
			Role:        string(user.Role),
			CreditLimit: user.CreditLimit,
			CurrentDue:  user.CurrentDue,
		})
	}

	return response, nil
}


func GetUserByID(id int32) (UserResponse, error) {

	user, err := config.Queries.GetUserByID(context.Background(), id)

	if err != nil {
		return UserResponse{}, err
	}


	response := UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Role:        string(user.Role),
		CreditLimit: user.CreditLimit,
		CurrentDue:  user.CurrentDue,
	}


	return response, nil
}


func UpdateUser(user generated.UpdateUserParams) error {

	err := config.Queries.UpdateUser(context.Background(), user)

	if err != nil {
		return err
	}

	return nil
}


func DeleteUser(id int32) error {

	err := config.Queries.DeleteUser(context.Background(), id)

	if err != nil {
		return err
	}

	return nil
}