package http

import (
	"log"
	"net/http"
)

func InitServer() {
	mux := http.NewServeMux()

	// Ingredient routes
	mux.HandleFunc("GET /api/ingredient", GetAllIngredients)

	// Cocktail routes
	mux.HandleFunc("GET /api/cocktail", GetAllCocktails)
	mux.HandleFunc("POST /api/cocktail", CreateCocktail)

	log.Print("Listening...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
