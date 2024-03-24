package ingredients

import (
	"context"
	"github.com/oli4maes/sipsavy/internal/infrastructure/mediator"
	"github.com/oli4maes/sipsavy/internal/infrastructure/persistence/relational"
	"log"
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
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GetAllIngredientsHandler interface {
	Handle() (GetAllIngredientsResponse, error)
}

// getAllIngredientsHandler is the medaitor handler, all dependencies should be added here
type getAllIngredientsHandler struct {
	repo relational.IngredientRepository
}

func (h getAllIngredientsHandler) Handle(request GetAllIngredientsRequest) (GetAllIngredientsResponse, error) {
	ingredients, err := h.repo.GetAll()
	if err != nil {
		log.Fatalf("could not fetch ingredients: %v", err)
	}

	var dtos []ingredientDto

	for _, i := range ingredients {
		dto := ingredientDto{
			Id:   i.Id,
			Name: i.Name,
		}

		dtos = append(dtos, dto)
	}

	response := GetAllIngredientsResponse{
		Ingredients: dtos,
	}

	return response, nil
}
