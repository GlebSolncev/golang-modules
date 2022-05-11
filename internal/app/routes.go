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
			TodoController := controllers.TodoController{HttpType: "api"}

			apiTodo.GET("", TodoController.Index)
			apiTodo.POST("", TodoController.Store)
			apiTodo.GET("/:id", TodoController.Show)
			apiTodo.PUT("/:id", TodoController.Update)
			apiTodo.DELETE("/:id", TodoController.Delete)
		}
	}

	rWeb := r.Group("/web")
	{
		wt := rWeb.Group("/todo")
		{
			TodoController := controllers.TodoController{HttpType: "web"}

			wt.GET("", TodoController.Index)
			wt.POST("/store", TodoController.Store)
			wt.GET("/:id", TodoController.Show)
			wt.POST("/:id", TodoController.Update)
			wt.GET("/:id/delete", TodoController.Delete)
		}
	}

}
