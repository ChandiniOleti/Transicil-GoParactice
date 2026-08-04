package config

import "userservice/generated"

var Queries *generated.Queries

func InitSQLC() {
	Queries = generated.New(DB)
}
