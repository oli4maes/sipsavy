package getcocktailsbyingredients

import "github.com/google/uuid"

type Response struct {
	Cocktails []cocktailDto `json:"cocktails"`
}

type cocktailDto struct {
	CocktailId uuid.UUID `json:"id"`
	Name       string    `json:"name"`
}
