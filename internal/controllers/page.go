package controllers

import (
	"crud/internal/models/todo"
	"crud/pkg/helpers"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type ResponseBody struct {
	Page  string
	Name  string
	Todos []todo.Attributes
	Todo  todo.Attributes
}

func HomePage(c echo.Context) error {

	return c.Render(http.StatusOK, "welcome.tmpl", ResponseBody{
		Page: "welcome",
	})
}

func GetTodos(c echo.Context) error {

	return c.Render(http.StatusOK, "todo.tmpl", ResponseBody{
		Page:  "index",
		Todos: todo.GetAll(),
		Todo:  todo.Attributes{},
	})
}

func ShowTodo(c echo.Context) error {
	var (
		id, err    = strconv.Atoi(c.Param("id"))
		statusCode = http.StatusOK
		model      = todo.FindById(id)
	)

	helpers.Check(err)

	return c.Render(statusCode, "todo.tmpl", ResponseBody{
		Name:  "show",
		Todos: todo.GetAll(),
		Todo:  model,
	})
}

func StoreTodo(c echo.Context) error {
	var (
		model = new(todo.Attributes)
		err   = c.Bind(model)
	)

	helpers.Check(err)
	todo.Store(model)

	return c.Redirect(http.StatusFound, "/todo")
}

func UpdateTodo(c echo.Context) error {
	var (
		model      = todo.SetModel(c)
		statusCode = http.StatusFound
	)

	if model.Slug == "" {
		statusCode = http.StatusNotFound
	}

	return c.Redirect(statusCode, "/todo/"+c.Param("id"))
}

func DeleteTodo(c echo.Context) error {
	var (
		id, _ = strconv.Atoi(c.Param("id"))
	)

	todo.DelModel(id)

	return c.Redirect(http.StatusFound, "/todo")
}
