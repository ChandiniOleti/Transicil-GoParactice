package crud

import (
	"context"
	"fmt"

	"mongodbgocrud/models"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func InsertEmployee(client *mongo.Client) {

	collection := client.Database("company").Collection("employees")

	employee := models.Employee{
		Name:       "Chandini",
		Age:        21,
		Department: "IT",
		Salary:     50000,
		City:       "Guntur",
	}

	result, err := collection.InsertOne(context.Background(), employee)

	if err != nil {
		fmt.Println("Insert Error:", err)
		return
	}

	fmt.Println("Employee Inserted Successfully")
	fmt.Println("Inserted ID:", result.InsertedID)
}