package services

import (
	"context"

	"paylaterservice/config"
	"paylaterservice/generated"
)

func CreateUser(user generated.CreateUserParams) error {

	_, err := config.Queries.CreateUser(context.Background(), user)

	if err != nil {
		return err
	}

	return nil
}

func GetUsers() ([]generated.User, error) {

	users, err := config.Queries.GetUsers(context.Background())

	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id int32) (generated.User, error) {

	user, err := config.Queries.GetUserByID(context.Background(), id)

	if err != nil {
		return generated.User{}, err
	}

	return user, nil
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