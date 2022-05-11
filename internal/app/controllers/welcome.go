package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Welcome struct {
	Controllers
	HttpType string
}

func (w Welcome) Index(c echo.Context) error {
	if w.HttpType == "api" {
		return c.JSON(http.StatusOK, Response{
			NamePage: "welcome",
		})
	}

	return c.Render(http.StatusOK, "welcome.tmpl", nil)
}
