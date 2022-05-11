package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang-modules/pkg/path"
	"os"
)

// Start work echo framework and setup settings.
// Call Routes for route list
// Use default middleware
// Setup host and port from env
func Start() {
	_ = godotenv.Load(".env")
	e := echo.New()
	e.Debug = true
	e.Renderer = templates
	Routes(e)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//p := prometheus.NewPrometheus("echo", nil)
	//p.Use(e)

	//host:port
	e.Logger.Fatal(e.Start(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT")))
}

// getDataSourceName call from Routes for set connect to DB
// Use .env for setup connection to DB (Def Sqlite)
func getDataSourceName() string {
	host := os.Getenv("DB_HOST")

	if os.Getenv("DB_DRIVER") == "file" {
		host = path.GetBasePath(host)
	}

	return fmt.Sprintf("%s:%s%s%s?%s",
		os.Getenv("DB_DRIVER"),
		host,
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDITION"))
}
