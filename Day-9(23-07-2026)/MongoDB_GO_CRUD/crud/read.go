package crud

import (
	"context"
	"fmt"

	"mongodbgocrud/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetAllEmployees(client *mongo.Client) {

	collection := client.Database("company").Collection("employees")

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		fmt.Println("Find Error:", err)
		return
	}

	defer cursor.Close(context.Background())

	var employees []models.Employee

	err = cursor.All(context.Background(), &employees)

	if err != nil {
		fmt.Println("Decode Error:", err)
		return
	}

	for _, emp := range employees {
		fmt.Println(emp)
	}
}

//=====findone
func GetEmployeeByName(client *mongo.Client, name string) {

	collection := client.Database("company").Collection("employees")

	var employee models.Employee

	err := collection.FindOne(
		context.Background(),
		bson.M{
			"name": name,
		},
	).Decode(&employee)

	if err != nil {
		fmt.Println("Employee Not Found:", err)
		return
	}

	fmt.Println(employee)
}