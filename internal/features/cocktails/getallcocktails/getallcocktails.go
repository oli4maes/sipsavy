package getallcocktails

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

	repo := relational.NewCocktailRepository(connString)

	err := mediator.Register[Request, Response](handler{repo})
	if err != nil {
		panic(err)
	}
}

type Request struct{}

type Response struct {
	Cocktails []cocktailDto `json:"cocktails"`
}

type cocktailDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Handler interface {
	Handle() (Response, error)
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
			Id:   c.Id,
			Name: c.Name,
		}

		dtos = append(dtos, dto)
	}

	response := Response{
		Cocktails: dtos,
	}

	return response, nil
}
