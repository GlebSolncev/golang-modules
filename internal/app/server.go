package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func Start() {
	e := echo.New()
	e.Debug = true
	e.Renderer = templates

	Routes(e)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//host:port
	e.Logger.Fatal(e.Start(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT")))
}
