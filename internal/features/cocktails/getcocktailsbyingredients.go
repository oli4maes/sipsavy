package cocktailfeatures

import (
	"context"
	"log"
	"os"

	"github.com/oli4maes/sipsavy/internal/infrastructure/mediator"
	"github.com/oli4maes/sipsavy/internal/infrastructure/persistence/relational"
)

// Register getCocktailsByIngredientsHandler
func init() {
	connString, exists := os.LookupEnv("CONNECTION_STRING")
	if !exists {
		panic("connection string env variable not set")
	}

	repo := relational.NewCocktailRepository(connString)

	err := mediator.Register[GetCocktailsByIngredientsRequest, GetCocktailsByIngredientsResponse](getCocktailsByIngredientsHandler{repo})
	if err != nil {
		panic(err)
	}
}

type GetCocktailsByIngredientsRequest struct {
	IngredientIds []int `json:"ingredient_ids"`
}

type GetCocktailsByIngredientsResponse struct {
	Cocktails []getCocktailsByIngredientsDto `json:"cocktails"`
}

type getCocktailsByIngredientsDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GetCocktailsByIngredientsHandler interface {
	Handle() (GetCocktailsByIngredientsResponse, error)
}

// getCocktailsByIngredientsHandler is the mediator handler, all dependencies should be added here
type getCocktailsByIngredientsHandler struct {
	repo relational.CocktailRepository
}

func (h getCocktailsByIngredientsHandler) Handle(ctx context.Context, request GetCocktailsByIngredientsRequest) (GetCocktailsByIngredientsResponse, error) {
	cocktails, err := h.repo.GetByIngredientIds(ctx, request.IngredientIds)
	if err != nil {
		log.Fatalf("could not fetch cocktails by ingredient ids: %v", err)
	}
	if cocktails == nil {
		return GetCocktailsByIngredientsResponse{Cocktails: []getCocktailsByIngredientsDto{}}, nil
	}

	var dtos []getCocktailsByIngredientsDto

	for _, c := range cocktails {
		dto := getCocktailsByIngredientsDto{
			Id:   c.Id,
			Name: c.Name,
		}

		dtos = append(dtos, dto)
	}

	response := GetCocktailsByIngredientsResponse{Cocktails: dtos}
	return response, nil
}
