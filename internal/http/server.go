package http

import (
	"log"
	"net/http"

	"github.com/oli4maes/sipsavy/internal/features/cocktails/createcocktail"
	"github.com/oli4maes/sipsavy/internal/features/cocktails/getallcocktails"
	"github.com/oli4maes/sipsavy/internal/features/ingredients/createingredient"
)

func InitServer() error {
	mux := http.NewServeMux()

	// Ingredient routes
	mux.HandleFunc("GET /api/ingredient", GetAllIngredients)
	mux.HandleFunc("POST /api/ingredient", createingredient.Handle)

	// Cocktail routes
	mux.HandleFunc("GET /api/cocktail", getallcocktails.Handle)
	mux.HandleFunc("POST /api/cocktail", createcocktail.Handle)

	log.Print("Listening...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return err
	}

	return nil
}
