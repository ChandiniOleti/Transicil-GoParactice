package config

import "paylaterservice/generated"

var Queries *generated.Queries

func InitSQLC() {
	Queries = generated.New(DB)
}