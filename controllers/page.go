package controllers

import (
	models "crud/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetTodos(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name":  "Index",
		"todos": models.GetAll(),
		"todo":  nil,
	})
}

func ShowTodo(c echo.Context) error {
	var (
		id, _      = strconv.Atoi(c.Param("id"))
		statusCode = http.StatusOK
		todo       = models.FindById(id)
	)

	if todo.Slug == "" {
		statusCode = http.StatusNotFound
		todo = models.Todo{}
	}

	return c.Render(statusCode, "home.html", map[string]interface{}{
		"name":  "Show",
		"todos": models.GetAll(),
		"todo":  todo,
	})
}

func StoreTodo(c echo.Context) error {
	models.StoreTodo(c)

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name":  "Store",
		"todos": models.GetAll(),
		"todo":  nil,
	})
}

func UpdateTodo(c echo.Context) error {
	var (
		todo       = models.UpdateTodo(c)
		statusCode = http.StatusOK
	)

	if todo.Slug == "" {
		statusCode = http.StatusNotFound
	}

	return c.Render(statusCode, "home.html", map[string]interface{}{
		"name":  "Update",
		"todos": models.GetAll(),
		"todo":  todo,
	})
}

func DeleteTodo(c echo.Context) error {
	var (
		statusCode = http.StatusOK
	)
	models.DeleteTodo(c)

	return c.Render(statusCode, "home.html", map[string]interface{}{
		"name":  "Home",
		"todos": models.GetAll(),
		"todo":  nil,
	})
}
