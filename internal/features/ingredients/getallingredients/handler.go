package getallingredients

import (
	"context"
	"log"

	"github.com/oli4maes/mediator"

	"github.com/oli4maes/sipsavy/internal/infrastructure/persistence/relational"
)

// Register handler
func init() {
	repo := relational.NewIngredientRepository()

	err := mediator.Register[Request, Response](handler{repo: repo})
	if err != nil {
		panic(err)
	}
}

// handler is the medaitor handler, all dependencies should be added here
type handler struct {
	repo relational.IngredientRepository
}

func (h handler) Handle(ctx context.Context, request Request) (Response, error) {
	ingredients, err := h.repo.GetAll(ctx)
	if err != nil {
		log.Fatalf("could not fetch ingredients: %v", err)
	}
	if ingredients == nil {
		return Response{Ingredients: []ingredientDto{}}, nil
	}

	var dtos []ingredientDto

	for _, i := range ingredients {
		dto := ingredientDto{
			IngredientId: i.Id,
			Name:         i.Name,
		}

		dtos = append(dtos, dto)
	}

	response := Response{
		Ingredients: dtos,
	}

	return response, nil
}
