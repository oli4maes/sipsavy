package cocktails

import (
	"context"
	"github.com/google/uuid"
	"github.com/oli4maes/sipsavy/internal/mediator"
)

// Register getAllCocktailsHandler
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

// getAllCocktailsHandler is the medaitor handler, all dependencies should be added here
type getAllCocktailsHandler struct{}

func (h getAllCocktailsHandler) Handle(request GetAllCocktailsRequest) (GetAllCocktailsResponse, error) {
	// TODO: fetch this data from a repository or a query facade?
	var cocktails []CocktailDto

	response := GetAllCocktailsResponse{
		Cocktails: cocktails,
	}

	return response, nil
}
