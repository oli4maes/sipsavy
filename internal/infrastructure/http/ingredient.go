package http

import (
	"encoding/json"
	"fmt"
	features "github.com/oli4maes/sipsavy/internal/features/ingredients"
	"github.com/oli4maes/sipsavy/internal/infrastructure/mediator"
	"net/http"
)

func GetAllIngredients(w http.ResponseWriter, r *http.Request) {
	req := features.GetAllIngredientsRequest{}
	res, err := mediator.Send[features.GetAllIngredientsRequest, features.GetAllIngredientsResponse](req)
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
	var req features.CreateIngredientRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := mediator.Send[features.CreateIngredientRequest, features.CreateIngredientResponse](req)
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
