package getallingredients

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oli4maes/mediator"
)

func Handle(c echo.Context) error {
	req := Request{}
	res, err := mediator.Send[Request, Response](context.Background(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
