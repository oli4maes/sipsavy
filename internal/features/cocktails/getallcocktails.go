package features

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

	repo := relational.NewCocktailRepository(connString, context.Background())

	err := mediator.Register[GetAllCocktailsRequest, GetAllCocktailsResponse](getAllCocktailsHandler{repo: repo})
	if err != nil {
		panic(err)
	}
}

type GetAllCocktailsRequest struct{}

type GetAllCocktailsResponse struct {
	Cocktails []getAllCocktailDto `json:"cocktails"`
}

type getAllCocktailDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GetAllCocktailsHandler interface {
	Handle() (GetAllCocktailsResponse, error)
}

// getAllCocktailsHandler is the medaitor handler, all dependencies should be added here
type getAllCocktailsHandler struct {
	repo relational.CocktailRepository
}

func (h getAllCocktailsHandler) Handle(request GetAllCocktailsRequest) (GetAllCocktailsResponse, error) {
	cocktails, err := h.repo.GetAll()
	if err != nil {
		log.Fatalf("could not fetch ingredients: %v", err)
	}
	if cocktails == nil {
		return GetAllCocktailsResponse{Cocktails: []getAllCocktailDto{}}, nil
	}

	var dtos []getAllCocktailDto

	for _, i := range cocktails {
		dto := getAllCocktailDto{
			Id:   i.Id,
			Name: i.Name,
		}

		dtos = append(dtos, dto)
	}

	response := GetAllCocktailsResponse{
		Cocktails: dtos,
	}

	return response, nil
}
