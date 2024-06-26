package createingredient

import (
	"context"
	"time"

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

// handler is the mediator handler, all dependencies should be added here
type handler struct {
	repo relational.IngredientRepository
}

func (h handler) Handle(ctx context.Context, request Request) (Response, error) {
	ingredient := relational.Ingredient{
		Name:           request.Name,
		Created:        time.Now(),
		CreatedBy:      "test",
		LastModified:   time.Now(),
		LastModifiedBy: "test",
	}

	createdIngredient, err := h.repo.Create(ctx, ingredient)
	if err != nil {
		return Response{}, err
	}

	return Response{
		Ingredient: ingredientDto{
			IngredientId: createdIngredient.Id,
			Name:         createdIngredient.Name,
		},
	}, nil
}
