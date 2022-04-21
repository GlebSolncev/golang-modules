package app

import (
	"crud/pkg/path"
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
		templates: template.Must(template.ParseGlob(path.GetBasePath("web/*.tmpl"))),
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
