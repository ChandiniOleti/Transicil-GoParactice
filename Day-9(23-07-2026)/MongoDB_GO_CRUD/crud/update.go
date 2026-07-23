package crud

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func UpdateEmployeeSalary(client *mongo.Client, name string, salary float64) {

	collection := client.Database("company").Collection("employees")

	filter := bson.M{
		"name": name,
	}

	update := bson.M{
		"$set": bson.M{
			"salary": salary,
		},
	}

	result, err := collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	if err != nil {
		fmt.Println("Update Error:", err)
		return
	}

	fmt.Println("Matched Count :", result.MatchedCount)
	fmt.Println("Modified Count:", result.ModifiedCount)
}