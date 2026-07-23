package crud

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func DeleteEmployee(client *mongo.Client, name string) {

	collection := client.Database("company").Collection("employees")

	filter := bson.M{
		"name": name,
	}

	result, err := collection.DeleteOne(
		context.Background(),
		filter,
	)

	if err != nil {
		fmt.Println("Delete Error:", err)
		return
	}

	fmt.Println("Deleted Count:", result.DeletedCount)
}