package app

import (
	"crud/controllers"
	"github.com/labstack/echo"
	"net/http"
)

func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	e.GET("/pages", controllers.GetPages)
	e.POST("/pages", controllers.StorePage)
	e.GET("/pages/:id", controllers.ShowPage)
	e.PUT("/pages/:id", controllers.UpdatePage)
}
