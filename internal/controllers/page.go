package controllers

import (
	"crud/internal/models"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type ResponseBody struct {
	Page  string
	Name  string
	Todos []models.Todo
	Todo  models.Todo
}

func HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "welcome.tmpl", ResponseBody{
		Page: "welcome",
	})
}

func GetTodos(c echo.Context) error {
	println("GET TODOS")
	return c.Render(http.StatusOK, "todo.tmpl", ResponseBody{
		Page:  "index",
		Todos: models.GetAllTodos(),
		Todo:  models.Todo{},
	})
}

func ShowTodo(c echo.Context) error {
	var (
		id, _      = strconv.Atoi(c.Param("id"))
		statusCode = http.StatusOK
		todo       = models.FindById(id)
	)

	//map[string]interface{}{
	return c.Render(statusCode, "todo.tmpl", ResponseBody{
		Name:  "show",
		Todos: models.GetAllTodos(),
		Todo:  todo,
	})
}

func StoreTodo(c echo.Context) error {
	models.StoreTodo(c)
	fmt.Println(1213123)

	return c.Redirect(http.StatusFound, "/todo")
}

func UpdateTodo(c echo.Context) error {
	var (
		todo       = models.UpdateTodo(c)
		statusCode = http.StatusFound
	)

	if todo.Slug == "" {
		statusCode = http.StatusNotFound
	}

	return c.Redirect(statusCode, "/todo/"+c.Param("id"))
}

func DeleteTodo(c echo.Context) error {
	models.DeleteTodo(c)

	return c.Redirect(http.StatusFound, "/todo")
}
