package app

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"os"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var (
	t = &Template{
		templates: template.Must(template.ParseGlob("web/*.tmpl")),
	}
)

func Start() {
	e := echo.New()
	e.Debug = true
	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	Routes(e)

	//host:port
	fmt.Println(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"))
	e.Logger.Fatal(e.Start(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT")))
}
