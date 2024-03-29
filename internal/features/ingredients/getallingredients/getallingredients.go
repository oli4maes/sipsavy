package getallingredients

import (
	"context"
	"log"
	"os"

	"github.com/oli4maes/sipsavy/internal/infrastructure/mediator"
	"github.com/oli4maes/sipsavy/internal/infrastructure/persistence/relational"
)

// Register handler
func init() {
	connString, exists := os.LookupEnv("CONNECTION_STRING")
	if !exists {
		panic("connection string env variable not set")
	}

	repo := relational.NewIngredientRepository(connString)

	err := mediator.Register[Request, Response](handler{repo: repo})
	if err != nil {
		panic(err)
	}
}

type Request struct{}

type Response struct {
	Ingredients []ingredientDto `json:"ingredients"`
}

type ingredientDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Handler interface {
	Handle() (Response, error)
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
			Id:   i.Id,
			Name: i.Name,
		}

		dtos = append(dtos, dto)
	}

	response := Response{
		Ingredients: dtos,
	}

	return response, nil
}