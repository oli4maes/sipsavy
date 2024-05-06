package createcocktail

import "github.com/google/uuid"

type Response struct {
	Cocktail cocktailDto `json:"cocktail"`
}

type cocktailDto struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
