package app

import (
	_ "embed"
	"github.com/GlebSolncev/golang-modules/internal/app/controllers"
	"github.com/GlebSolncev/golang-modules/pkg/template"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func Routes(e *echo.Echo) {

	/** SWAGGER **/

	e.GET("swag/*", echoSwagger.WrapHandler)

	/** FS **/

	staticHandler := http.FileServer(template.GetFileSystem())
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", staticHandler)))

	assetHandler := http.FileServer(template.GetAssets())
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", assetHandler)))

	e.GET("/", controllers.Welcome{}.Index)

	/** STATUS **/

	apiS := e.Group("/api")
	ApiStatus := apiS.Group("/status")
	ApiStatus.GET("", controllers.StatusController{}.Index)
	ApiStatus.POST("", controllers.StatusController{}.Store)
	ApiStatus.GET("/:id", controllers.StatusController{}.Show)
	ApiStatus.PUT("/:id", controllers.StatusController{}.Update)
	ApiStatus.DELETE("/:id", controllers.StatusController{}.Delete)

	/** TODO **/

	apiT := e.Group("/api")
	ApiTodo := apiT.Group("/todo")
	ApiTodo.GET("", controllers.TodoController{HttpType: "api"}.Index)
	ApiTodo.POST("", controllers.TodoController{HttpType: "api"}.Store)
	ApiTodo.GET("/:id", controllers.TodoController{HttpType: "api"}.Show)
	ApiTodo.PUT("/:id", controllers.TodoController{HttpType: "api"}.Update)
	ApiTodo.DELETE("/:id", controllers.TodoController{HttpType: "api"}.Delete)

	web := e.Group("/web")
	WebTodo := web.Group("/todo")
	WebTodo.GET("", controllers.TodoController{HttpType: "web"}.Index)
	WebTodo.POST("/store", controllers.TodoController{HttpType: "web"}.Store)
	WebTodo.GET("/:id", controllers.TodoController{HttpType: "web"}.Show)
	WebTodo.POST("/:id", controllers.TodoController{HttpType: "web"}.Update)
	WebTodo.GET("/:id/delete", controllers.TodoController{HttpType: "web"}.Delete)

}
