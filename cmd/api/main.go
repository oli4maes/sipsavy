package main

import (
	"github.com/labstack/echo/v4"
	"github.com/oli4maes/sipsavy/handler"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)

	app.Start(":3000")

	//err := relational.MigrateDb()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = http.InitServer()
	//if err != nil {
	//	fmt.Println(err)
	//}
}
