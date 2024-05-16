package getallcocktails

import (
	"context"
	"log"

	"github.com/oli4maes/mediator"

	"github.com/oli4maes/sipsavy/internal/infrastructure/persistence/relational"
)

// Register handler
func init() {
	repo := relational.NewCocktailRepository()

	err := mediator.Register[Request, Response](handler{repo})
	if err != nil {
		panic(err)
	}
}

// handler is the medaitor handler, all dependencies should be added here
type handler struct {
	repo relational.CocktailRepository
}

func (h handler) Handle(ctx context.Context, request Request) (Response, error) {
	cocktails, err := h.repo.GetAll(ctx)
	if err != nil {
		log.Fatalf("could not fetch ingredients: %v", err)
	}
	if cocktails == nil {
		return Response{Cocktails: []cocktailDto{}}, nil
	}

	var dtos []cocktailDto

	for _, c := range cocktails {
		dto := cocktailDto{
			CocktailId: c.Id,
			Name:       c.Name,
		}

		dtos = append(dtos, dto)
	}

	response := Response{
		Cocktails: dtos,
	}

	return response, nil
}
