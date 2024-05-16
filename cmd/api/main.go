package main

import (
	"github.com/oli4maes/sipsavy/internal/http"
	"github.com/oli4maes/sipsavy/internal/infrastructure/persistence/relational"
)

func main() {
	err := relational.MigrateDb()
	if err != nil {
		panic(err)
	}

	http.InitServer()
}
