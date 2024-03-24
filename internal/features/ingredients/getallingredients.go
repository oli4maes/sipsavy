package ingredients

import (
	"context"
	"github.com/oli4maes/sipsavy/internal/infrastructure/mediator"
	"github.com/oli4maes/sipsavy/internal/infrastructure/persistence/relational"
	"os"
)

// Register getAllCocktailsHandler
func init() {
	connString, exists := os.LookupEnv("CONNECTION_STRING")
	if !exists {
		panic("connection string env variable not set")
	}

	repo := relational.NewIngredientRepository(connString)

	err := mediator.Register[GetAllIngredientsRequest, GetAllIngredientsResponse](getAllIngredientsHandler{repo: repo})
	if err != nil {
		panic(err)
	}
}

type GetAllIngredientsRequest struct {
	Ctx context.Context
}

type GetAllIngredientsResponse struct {
	Ingredients []ingredientDto
}

type ingredientDto struct {
	Id   int
	Name string
}

type GetAllIngredientsHandler interface {
	Handle() (GetAllIngredientsResponse, error)
}

// getAllIngredientsHandler is the medaitor handler, all dependencies should be added here
type getAllIngredientsHandler struct {
	repo relational.IngredientRepository
}

func (h getAllIngredientsHandler) Handle(request GetAllIngredientsRequest) (GetAllIngredientsResponse, error) {
	// TODO: fetch this data from a repository or a query facade?
	var ingredients []ingredientDto

	response := GetAllIngredientsResponse{
		Ingredients: ingredients,
	}

	return response, nil
}
