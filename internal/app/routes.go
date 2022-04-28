package app

import (
	"crud"
	"crud/internal/controllers"
	_ "embed"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Routes(e *echo.Echo) {
	staticHandler := http.FileServer(crud.GetFileSystem())
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", staticHandler)))

	assetHandler := http.FileServer(crud.GetAssets())
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", assetHandler)))

	e.GET("/", controllers.HomePage)

	api := e.Group("/api")
	ApiTodo := api.Group("/todo")
	ApiTodo.GET("", controllers.GetTodos)
	ApiTodo.POST("/store", controllers.StoreTodo)
	ApiTodo.GET("/:id", controllers.ShowTodo)
	ApiTodo.POST("/:id", controllers.UpdateTodo)
	ApiTodo.GET("/:id/delete", controllers.DeleteTodo)

	web := e.Group("/web")
	WebTodo := web.Group("/todo")
	WebTodo.GET("", controllers.GetTodos)
	WebTodo.POST("/store", controllers.StoreTodo)
	WebTodo.GET("/:id", controllers.ShowTodo)
	WebTodo.POST("/:id", controllers.UpdateTodo)
	WebTodo.GET("/:id/delete", controllers.DeleteTodo)

}
