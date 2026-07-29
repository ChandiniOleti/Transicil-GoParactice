package services

import (
	"context"
	"errors"

	"paylaterservice/config"
	"paylaterservice/utils"
)


type LoginRequest struct {

	Email string `json:"email"`

	Password string `json:"password"`
}



func Login(request LoginRequest) (string,error){


	ctx := context.Background()


	user,err := config.Queries.GetUserByEmail(ctx,request.Email)


	if err != nil {

		return "",errors.New("invalid email")
	}


	if user.Password != request.Password {

		return "",errors.New("invalid password")
	}


	token, err := utils.GenerateToken(
	user.ID,
	string(user.Role),
)

	if err != nil {

		return "",err
	}


	return token,nil
}