package config

import (
	"employeecrudsqlc/generated"
)

var Queries *generated.Queries

func InitQueries() {

	Queries = generated.New(DB)

}