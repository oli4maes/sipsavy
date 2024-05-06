package getallingredients

import "github.com/google/uuid"

type Response struct {
	Ingredients []ingredientDto `json:"ingredients"`
}

type ingredientDto struct {
	IngredientId uuid.UUID `json:"id"`
	Name         string    `json:"name"`
}
