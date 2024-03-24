package api

import (
	"context"
	"github.com/google/uuid"
	"sipsavy/internal/mediator"
)

func init() {
	err := mediator.Register[GetAllCocktailsRequest, GetAllCocktailsResponse](getAllCocktailsHandler{})
	if err != nil {
		panic(err)
	}
}

type GetAllCocktailsRequest struct {
	Ctx context.Context
}

type GetAllCocktailsResponse struct {
	Cocktails []CocktailDto
}

type CocktailDto struct {
	Id   uuid.UUID
	Name string
}

type GetAllCocktailsHandler interface {
	Handle() (GetAllCocktailsResponse, error)
}

type getAllCocktailsHandler struct{}

func (h getAllCocktailsHandler) Handle(request GetAllCocktailsRequest) (GetAllCocktailsResponse, error) {
	var cocktails []CocktailDto

	response := GetAllCocktailsResponse{
		Cocktails: cocktails,
	}

	return response, nil
}
