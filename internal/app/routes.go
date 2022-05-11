package app

import (
	_ "embed"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang-modules/internal/app/controllers"
	"golang-modules/pkg/template"
	"net/http"
)

func Routes(r *echo.Echo) {

	/** SWAGGER **/

	r.GET("swag/*", echoSwagger.WrapHandler)

	/** FS **/

	staticHandler := http.FileServer(template.GetFileSystem())
	r.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", staticHandler)))

	assetHandler := http.FileServer(template.GetAssets())
	r.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", assetHandler)))

	r.GET("/", controllers.Welcome{}.Index)

	// Login route
	r.POST("/auth", controllers.Auth)

	/** STATUS **/
	rApi := r.Group("/api")
	{
		apiStatus := rApi.Group("/status")
		{
			apiStatus.GET("", controllers.StatusController{}.Index)
		}

		/** TODOs **/

		config := middleware.JWTConfig{
			Claims:     &controllers.JwtCustomClaims{},
			SigningKey: []byte("secret"),
		}
		apiTodo := rApi.Group("/todo")
		apiTodo.Use(middleware.JWTWithConfig(config))
		{
			apiTodo.GET("", controllers.TodoController{HttpType: "api"}.Index)
			apiTodo.POST("", controllers.TodoController{HttpType: "api"}.Store)
			apiTodo.GET("/:id", controllers.TodoController{HttpType: "api"}.Show)
			apiTodo.PUT("/:id", controllers.TodoController{HttpType: "api"}.Update)
			apiTodo.DELETE("/:id", controllers.TodoController{HttpType: "api"}.Delete)
		}
	}

	rWeb := r.Group("/web")
	{
		wt := rWeb.Group("/todo")
		{
			wt.GET("", controllers.TodoController{HttpType: "web"}.Index)
			wt.POST("/store", controllers.TodoController{HttpType: "web"}.Store)
			wt.GET("/:id", controllers.TodoController{HttpType: "web"}.Show)
			wt.POST("/:id", controllers.TodoController{HttpType: "web"}.Update)
			wt.GET("/:id/delete", controllers.TodoController{HttpType: "web"}.Delete)
		}
	}

}
