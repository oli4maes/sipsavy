package http

import (
	"github.com/labstack/echo/v4"
	"github.com/oli4maes/sipsavy/internal/features/cocktails/createcocktail"
	"github.com/oli4maes/sipsavy/internal/features/cocktails/getallcocktails"
	"github.com/oli4maes/sipsavy/internal/features/ingredients/createingredient"
	"github.com/oli4maes/sipsavy/internal/features/ingredients/getallingredients"
)

func InitServer() {
	e := echo.New()

	// Ingredient endpoints
	e.GET("/api/ingredient", getallingredients.Handle)
	e.POST("/api/ingredient", createingredient.Handle)

	// Cocktail endpoints
	e.GET("/api/cocktail", getallcocktails.Handle)
	e.POST("/api/cocktail", createcocktail.Handle)

	e.Logger.Fatal(e.Start(":8080"))
}
