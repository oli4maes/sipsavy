package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	features "github.com/oli4maes/sipsavy/internal/features/ingredients"
	"github.com/oli4maes/sipsavy/internal/infrastructure/mediator"
)

func GetAllIngredients(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := features.GetAllIngredientsRequest{}
	res, err := mediator.Send[features.GetAllIngredientsRequest, features.GetAllIngredientsResponse](req, ctx)
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

func CreateIngredient(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req features.CreateIngredientRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := mediator.Send[features.CreateIngredientRequest, features.CreateIngredientResponse](req, ctx)
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
