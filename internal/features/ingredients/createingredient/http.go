package createingredient

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oli4maes/mediator"
)

func Handle(c echo.Context) error {
	req := new(Request)
	if err := c.Bind(req); err != nil {
		return err
	}

	err := req.Validate()
	if err != nil {
		return err
	}

	res, err := mediator.Send[Request, Response](context.Background(), *req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}
