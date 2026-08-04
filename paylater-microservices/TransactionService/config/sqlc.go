package config

import "transactionservice/generated"

var Queries *generated.Queries

func InitSQLC() {
	Queries = generated.New(DB)
}
