package controllers

import "github.com/labstack/echo/v4"

type (
	Response struct {
		Page    string
		Payload interface{}
	}

	Controllers interface {
		Index(c echo.Context) error
		Store(c echo.Context) error
		Show(c echo.Context) error
		Update(c echo.Context) error
		Delete(c echo.Context) error
	}
)
