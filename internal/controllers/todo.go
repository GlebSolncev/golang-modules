package controllers

import (
	"crud/internal/models/todo"
	"crud/pkg/ent"
	model "crud/pkg/ent/todo"
	"crud/pkg/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	_ "unsafe"
)

type (
	Payload struct {
		Page   string
		Status []model.Status
		Todos  []ent.Todo
		Todo   ent.Todo
	}
)

var payload Payload

func HomePage(c echo.Context) error {
	payload = Payload{
		Page: "welcome",
	}

	if strings.Contains(c.Path(), "api") {
		return c.JSON(http.StatusOK, payload)
	} else {
		return c.Render(http.StatusOK, "welcome.tmpl", Payload{
			Page: "welcome",
		})
	}
}

func GetTodos(c echo.Context) error {
	payload = Payload{
		Page:   "index",
		Todos:  todo.GetAll(),
		Status: todo.GetStatuses(),
	}

	if strings.Contains(c.Path(), "api") {
		return c.JSON(http.StatusOK, payload)
	} else {
		return c.Render(http.StatusOK, "todo.tmpl", payload)
	}
}

func ShowTodo(c echo.Context) error {
	var (
		id, err = strconv.Atoi(c.Param("id"))
		m       = todo.FindById(id)
	)

	helpers.Check(err)
	payload = Payload{
		Page: "Show",
		Todo: m,
	}

	if strings.Contains(c.Path(), "api") {
		return c.JSON(http.StatusOK, payload)
	} else {
		return c.Render(http.StatusOK, "todo.tmpl", payload)
	}
}

func StoreTodo(c echo.Context) error {
	var (
		m = new(ent.Todo)
	)
	if err := c.Bind(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	todo.Store(m)

	if strings.Contains(c.Path(), "api") {
		return c.JSON(http.StatusOK, todo.GetAll())
	} else {
		return c.Redirect(http.StatusFound, "/web/todo")
	}

}

func UpdateTodo(c echo.Context) error {
	var (
		id  = c.Param("id")
		m   = new(ent.Todo)
		err error
	)

	if err = c.Bind(m); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	m.ID, err = strconv.Atoi(id)
	helpers.Check(err)
	m = todo.SetModel(m)

	if strings.Contains(c.Path(), "api") {
		return c.JSON(http.StatusOK, m)
	} else {
		return c.Redirect(http.StatusFound, "/web/todo/"+c.Param("id"))
	}
}

func DeleteTodo(c echo.Context) error {
	var (
		id, _ = strconv.Atoi(c.Param("id"))
	)

	todo.DelModel(id)

	if strings.Contains(c.Path(), "api") {
		return c.JSON(http.StatusOK, "OK")
	}
	return c.Redirect(http.StatusFound, "/web/todo/")
}
