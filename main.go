package main

import (
	"moduit/app/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/one", handler.One)
	e.GET("/two", handler.Two)
	e.GET("/three", handler.Three)
	e.Logger.Fatal(e.Start(":1323"))
}
