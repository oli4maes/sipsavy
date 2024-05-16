package createcocktail

import (
	"context"
	"time"

	"github.com/oli4maes/mediator"

	"github.com/oli4maes/sipsavy/internal/infrastructure/persistence/relational"
)

// Register createCocktailHandler
func init() {
	repo := relational.NewCocktailRepository()

	err := mediator.Register[Request, Response](handler{repo: repo})
	if err != nil {
		panic(err)
	}
}

// handler is the mediator handler, all dependencies should be added here
type handler struct {
	repo relational.CocktailRepository
}

func (h handler) Handle(ctx context.Context, request Request) (Response, error) {
	var ingredients []relational.CocktailIngredient

	for _, i := range request.Ingredients {
		ingredient := relational.CocktailIngredient{
			IngredientId:   i.IngredientId,
			Amount:         i.Amount,
			IngredientUnit: i.Unit,
		}

		ingredients = append(ingredients, ingredient)
	}

	cocktail := relational.Cocktail{
		Name:           request.Name,
		Created:        time.Now(),
		CreatedBy:      "test",
		LastModified:   time.Now(),
		LastModifiedBy: "test",
	}

	createdCocktail, err := h.repo.Create(ctx, cocktail)
	if err != nil {
		return Response{}, err
	}

	return Response{
		Cocktail: cocktailDto{
			Id:   createdCocktail.Id,
			Name: createdCocktail.Name,
		},
	}, nil
}
