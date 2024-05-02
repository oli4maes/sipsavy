package createcocktail

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
)

type Request struct {
	Name        string                    `json:"name"`
	Ingredients []IngredientWithAmountDto `json:"ingredients"`
}

type IngredientWithAmountDto struct {
	IngredientId uuid.UUID `json:"ingredient_id"`
	Amount       int       `json:"amount"`
	Unit         string    `json:"unit"`
}

func (r Request) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.Length(3, 250)))
}

func (i IngredientWithAmountDto) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.IngredientId, validation.Required, validation.Min(1)),
		validation.Field(&i.Amount, validation.Required, validation.Min(0)),
		validation.Field(&i.Unit, validation.Required))
}
