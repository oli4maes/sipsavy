package features

import (
	"github.com/oli4maes/sipsavy/internal/infrastructure/mediator"
	"github.com/oli4maes/sipsavy/internal/infrastructure/persistence/relational"
	"os"
)

// Register createCocktailHandler
func init() {
	connString, exists := os.LookupEnv("CONNECTION_STRING")
	if !exists {
		panic("connection string env variable not set")
	}

	repo := relational.NewCocktailRepository(connString)

	err := mediator.Register[CreateCocktailRequest, CreateCocktailResponse](createCocktailHandler{repo: repo})
	if err != nil {
		panic(err)
	}
}

type CreateCocktailRequest struct {
	Name        string
	Ingredients []IngredientWithAmountDto
}

type IngredientWithAmountDto struct {
	IngredientId int
	Amount       int
	Unit         int
}

type CreateCocktailResponse struct {
	Cocktail createCocktailDto `json:"cocktail"`
}

type createCocktailDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCocktailHandler interface {
	Handle() (CreateCocktailRequest, error)
}

// createCocktailHandler is the mediator handler, all dependencies should be added here
type createCocktailHandler struct {
	repo relational.CocktailRepository
}

func (h createCocktailHandler) Handle(request CreateCocktailRequest) (CreateCocktailResponse, error) {
	return CreateCocktailResponse{}, nil
}
