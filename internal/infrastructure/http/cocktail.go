package http

import (
	"encoding/json"
	"fmt"
	"github.com/oli4maes/sipsavy/internal/features/cocktails"
	"github.com/oli4maes/sipsavy/internal/infrastructure/mediator"
	"net/http"
)

func GetAllCocktails(w http.ResponseWriter, r *http.Request) {
	req := cocktails.GetAllCocktailsRequest{}
	res, err := mediator.Send[cocktails.GetAllCocktailsRequest, cocktails.GetAllCocktailsResponse](req)
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
