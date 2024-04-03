package getcocktailsbyingredients

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Request struct {
	IngredientIds []int `json:"ingredient_ids"`
}

func (r Request) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.IngredientIds, validation.Required, validation.Each(is.Int)))
}
