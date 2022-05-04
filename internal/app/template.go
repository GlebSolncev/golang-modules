package app

import (
	"github.com/GlebSolncev/golang-modules/pkg/path"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

var (
	templates = &Template{
		templates: template.Must(template.ParseGlob(path.GetBasePath("/pkg/template/web/*.tmpl"))),
	}
)

func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
