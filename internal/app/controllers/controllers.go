package controllers

import (
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "golang-modules/docs" // swagger
	"net/http"
	"os"
	"time"
)

type (
	Response struct {
		NamePage string
		Payload  interface{}
	}

	JwtCustomClaims struct {
		Name  string `json:"name"`
		Admin bool   `json:"admin"`
		jwt.StandardClaims
	}

	Controllers interface {
		Index(c echo.Context) error
		Store(c echo.Context) error
		Show(c echo.Context) error
		Update(c echo.Context) error
		Delete(c echo.Context) error
	}
)

// Auth godoc
// @Summary Login for work with Tоdo list
// @Description Get token for work with Tоdo list
// @Tags Auth
// @Accept json
// @Produce json
// @Param   username  query     string     false  "Username for login. ex 'hello'" hello
// @Param   password  query    string     false  "Password for login. ex 'world'" world
// @Success 200 {object} Response
// @Router /auth [POST]
func Auth(c echo.Context) error {
	_ = godotenv.Load(".env")
	var (
		username = c.FormValue("username")
		password = c.FormValue("password")
	)

	// Throws unauthorized error
	if username != os.Getenv("AUTH_USERNAME") || password != os.Getenv("AUTH_PASSWORD") {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &JwtCustomClaims{
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
