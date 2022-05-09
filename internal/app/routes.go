package app

import (
	_ "embed"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang-modules/internal/app/controllers"
	"golang-modules/pkg/template"
	"net/http"
	"time"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	fmt.Println(username, password)
	// Throws unauthorized error
	if username != "hello" || password != "world" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Admin",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

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
	r.POST("/login", login)

	/** STATUS **/
	rApi := r.Group("/api")
	{
		rs := rApi.Group("/status")
		{
			rs.GET("", controllers.StatusController{}.Index)
			rs.POST("", controllers.StatusController{}.Store)
			rs.GET("/:id", controllers.StatusController{}.Show)
			rs.PUT("/:id", controllers.StatusController{}.Update)
			rs.DELETE("/:id", controllers.StatusController{}.Delete)
		}

		/** TODOs **/

		config := middleware.JWTConfig{
			Claims:     &jwtCustomClaims{},
			SigningKey: []byte("secret"),
		}
		rs.Use(middleware.JWTWithConfig(config))
		ApiTodo := rApi.Group("/todo")
		{
			ApiTodo.GET("", controllers.TodoController{HttpType: "api"}.Index)
			ApiTodo.POST("", controllers.TodoController{HttpType: "api"}.Store)
			ApiTodo.GET("/:id", controllers.TodoController{HttpType: "api"}.Show)
			ApiTodo.PUT("/:id", controllers.TodoController{HttpType: "api"}.Update)
			ApiTodo.DELETE("/:id", controllers.TodoController{HttpType: "api"}.Delete)
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
