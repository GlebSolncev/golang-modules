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

	e.GET("/todos", controllers.GetTodos)
	e.POST("/todos", controllers.StoreTodo)
	e.GET("/todos/:id", controllers.ShowTodo)
	e.POST("/todos/:id", controllers.UpdateTodo)
	e.GET("/todos/:id/destroy", controllers.DeleteTodo)

}
