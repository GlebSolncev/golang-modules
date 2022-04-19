package app

import (
	"crud/internal/controllers"
	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	e.GET("/", controllers.HomePage)

	todo := e.Group("/todo")
	todo.GET("", controllers.GetTodos)
	todo.POST("/store", controllers.StoreTodo)
	todo.GET("/:id", controllers.ShowTodo)
	todo.POST("/:id", controllers.UpdateTodo)
	todo.GET("/:id/delete", controllers.DeleteTodo)

}
