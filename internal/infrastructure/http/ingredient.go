package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oli4maes/sipsavy/internal/features/ingredients/createingredient"
	"github.com/oli4maes/sipsavy/internal/features/ingredients/getallingredients"
	"github.com/oli4maes/sipsavy/internal/infrastructure/mediator"
)

func GetAllIngredients(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := getallingredients.Request{}
	res, err := mediator.Send[getallingredients.Request, getallingredients.Response](req, ctx)
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
	var req createingredient.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := mediator.Send[createingredient.Request, createingredient.Response](req, ctx)
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
