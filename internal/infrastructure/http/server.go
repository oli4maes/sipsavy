package http

import (
	"github.com/oli4maes/sipsavy/internal/infrastructure/http/ingredients"
	"log"
	"net/http"
)

func InitServer() {
	mux := http.NewServeMux()

	// Ingredient routes
	mux.HandleFunc("GET /api/ingredient", ingredients.GetAll)

	log.Print("Listening...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
