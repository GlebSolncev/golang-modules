package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var t = &Template{
	templates: template.Must(template.ParseGlob("resources/view/*.html")),
}

func Start() {
	e := echo.New()
	e.Renderer = t

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//todo rm
	e.GET("/hello", Hello)
	Routes(e)
	e.Logger.Fatal(e.Start(":8081"))
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", "World")
}
