package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// instanc echo
	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "hello")
	})
	e.Logger.Fatal(e.Start(":1234"))
}
