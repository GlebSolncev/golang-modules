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

	todo := e.Group("/todo")
	todo.GET("", controllers.GetTodos).Name = "index"
	todo.POST("/store", controllers.StoreTodo)
	todo.GET("/:id", controllers.ShowTodo).Name = "show"
	todo.POST("/:id", controllers.UpdateTodo)
	todo.GET("/:id/delete", controllers.DeleteTodo)

}
