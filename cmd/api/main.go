package main

import (
	"fmt"

	"github.com/oli4maes/sipsavy/internal/http"
	"github.com/oli4maes/sipsavy/internal/infrastructure/persistence/relational"
)

func main() {
	err := relational.MigrateDb()
	if err != nil {
		fmt.Println(err)
	}

	err = http.InitServer()
	if err != nil {
		fmt.Println(err)
	}
}
