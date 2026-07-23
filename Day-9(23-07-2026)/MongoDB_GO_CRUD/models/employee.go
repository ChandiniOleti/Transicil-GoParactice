package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Employee struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	Name       string        `bson:"name"`
	Age        int           `bson:"age"`
	Department string        `bson:"department"`
	Salary     float64       `bson:"salary"`
	City       string        `bson:"city"`
}