package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	Routes(e)
	e.Logger.Fatal(e.Start(":8081"))
}
