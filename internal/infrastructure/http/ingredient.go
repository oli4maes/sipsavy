package http

import (
	"encoding/json"
	"fmt"
	"github.com/oli4maes/sipsavy/internal/features/ingredients"
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