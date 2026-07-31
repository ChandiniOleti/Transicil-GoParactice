package services

import (
	"context"
	

	"paylaterservice/config"
	"paylaterservice/generated"
	

	"golang.org/x/crypto/bcrypt"
)



func GetAdmins() ([]generated.Admin, error) {

	admins, err := config.Queries.GetAdmins(context.Background())

	if err != nil {
		return nil, err
	}

	return admins, nil
}


func CreateAdmin(admin generated.CreateAdminParams) error {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(admin.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	admin.Password = string(hashedPassword)

	_, err = config.Queries.CreateAdmin(
		context.Background(),
		admin,
	)

	return err
}