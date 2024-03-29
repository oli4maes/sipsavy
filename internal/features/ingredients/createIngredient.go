package ingredientfeatures

import (
	"context"
	"os"
	"time"

	"github.com/oli4maes/sipsavy/internal/infrastructure/mediator"
	"github.com/oli4maes/sipsavy/internal/infrastructure/persistence/relational"
)

// Register createIngredientHandler
func init() {
	connString, exists := os.LookupEnv("CONNECTION_STRING")
	if !exists {
		panic("connection string env variable not set")
	}

	repo := relational.NewIngredientRepository(connString)

	err := mediator.Register[CreateIngredientRequest, CreateIngredientResponse](createIngredientHandler{repo: repo})
	if err != nil {
		panic(err)
	}
}

type CreateIngredientRequest struct {
	Name string `json:"name"`
}

type CreateIngredientResponse struct {
	Ingredient createIngredientDto `json:"ingredient"`
}

type createIngredientDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateIngredientHandler interface {
	Handle() (CreateIngredientRequest, error)
}

// createIngredientHandler is the mediator handler, all dependencies should be added here
type createIngredientHandler struct {
	repo relational.IngredientRepository
}

func (h createIngredientHandler) Handle(ctx context.Context, request CreateIngredientRequest) (CreateIngredientResponse, error) {
	ingredient := relational.Ingredient{
		Name:           request.Name,
		Created:        time.Now(),
		CreatedBy:      "test",
		LastModified:   time.Now(),
		LastModifiedBy: "test",
	}

	createdIngredient, err := h.repo.Create(ctx, ingredient)
	if err != nil {
		return CreateIngredientResponse{}, err
	}

	return CreateIngredientResponse{
		Ingredient: createIngredientDto{
			Id:   createdIngredient.Id,
			Name: createdIngredient.Name,
		},
	}, nil
}
