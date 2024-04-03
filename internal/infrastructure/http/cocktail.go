package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oli4maes/sipsavy/internal/features/cocktails/createcocktail"
	"github.com/oli4maes/sipsavy/internal/features/cocktails/getallcocktails"
	"github.com/oli4maes/sipsavy/internal/features/cocktails/getcocktailsbyingredients"
	"github.com/oli4maes/sipsavy/internal/infrastructure/mediator"
)

func GetAllCocktails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := getallcocktails.Request{}
	res, err := mediator.Send[getallcocktails.Request, getallcocktails.Response](ctx, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
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
	var req getcocktailsbyingredients.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = req.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := mediator.Send[getcocktailsbyingredients.Request, getcocktailsbyingredients.Response](ctx, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
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
	var req createcocktail.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = req.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := mediator.Send[createcocktail.Request, createcocktail.Response](ctx, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}
