package config

import "authservice/generated"

var Queries *generated.Queries

func InitSQLC() {
	Queries = generated.New(DB)
}
