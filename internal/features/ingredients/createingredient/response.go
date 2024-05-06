package createingredient

import "github.com/google/uuid"

type Response struct {
	Ingredient ingredientDto `json:"ingredient"`
}

type ingredientDto struct {
	IngredientId uuid.UUID `json:"id"`
	Name         string    `json:"name"`
}
