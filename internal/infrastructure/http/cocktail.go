package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	features "github.com/oli4maes/sipsavy/internal/features/cocktails"
	"github.com/oli4maes/sipsavy/internal/infrastructure/mediator"
)

func GetAllCocktails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := features.GetAllCocktailsRequest{}
	res, err := mediator.Send[features.GetAllCocktailsRequest, features.GetAllCocktailsResponse](req, ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}

func GetCocktailsByIngredientIds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req features.GetCocktailsByIngredientsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := mediator.Send[features.GetCocktailsByIngredientsRequest, features.GetCocktailsByIngredientsResponse](req, ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}

func CreateCocktail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req features.CreateCocktailRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := mediator.Send[features.CreateCocktailRequest, features.CreateCocktailResponse](req, ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}
