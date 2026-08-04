package config

import "merchantservice/generated"

var Queries *generated.Queries

func InitSQLC() {
	Queries = generated.New(DB)
}
