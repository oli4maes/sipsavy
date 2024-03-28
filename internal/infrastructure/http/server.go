package http

import (
	"log"
	"net/http"
)

func InitServer() {
	mux := http.NewServeMux()

	// Ingredient routes
	mux.HandleFunc("GET /api/ingredient", GetAllIngredients)
	mux.HandleFunc("POST /api/ingredient", CreateIngredient)

	// Cocktail routes
	mux.HandleFunc("GET /api/cocktail", GetAllCocktails)
	mux.HandleFunc("POST /api/cocktail", CreateCocktail)
	mux.HandleFunc("POST /api/cocktail/ingredients", GetCocktailsByIngredientIds)

	log.Print("Listening...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}
