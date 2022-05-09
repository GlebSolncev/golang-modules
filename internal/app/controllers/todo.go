package controllers

import (
	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger"   // echo-swagger middleware
	_ "golang-modules/internal/app/docs" // swagger
	"golang-modules/internal/app/models"
	"golang-modules/pkg/ent"
	"golang-modules/pkg/helpers"
	"net/http"
	"strconv"
	_ "unsafe"
)

type (
	TodoController struct {
		Controllers
		HttpType string
	}

	TodoS struct {
		// ID of the ent.
		ID int `json:"id"`
		// Slug holds the value of the "slug" field.
		Slug string `json:"slug"`
		// Name holds the value of the "name" field.
		Name *string `json:"name"`
		// Status holds the value of the "status" field.
		Status models.TodoStatus `json:"status"`
	}
)

var (
	todo = models.TodoModel{}
)

// Index godoc
// @Summary All todos items with statuses
// @Description Get todos list with status items
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /todo [get]
func (tc TodoController) Index(c echo.Context) error {
	res, err := todo.GetAll()
	helpers.Check(err)

	if tc.HttpType == "api" {
		return c.JSON(http.StatusOK, Response{Page: "Index", Payload: res})
	} else {
		return c.Render(http.StatusOK, "todo.tmpl", res)
	}
}

// Show godoc
// @Summary Item from todos list
// @Description Show item of Todos list
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 404
// @Router /todo/{ent.Todo.id} [get]
func (tc TodoController) Show(c echo.Context) error {
	var (
		id  int
		err error
		res interface{}
	)
	id, err = strconv.Atoi(c.Param("id"))
	helpers.Check(err)
	res, err = todo.FindById(id)
	helpers.Check(err)

	if tc.HttpType == "api" {
		return c.JSON(http.StatusOK, Response{Page: "Show", Payload: res})
	} else {
		return c.Render(http.StatusOK, "todo.tmpl", res)
	}
}

// Store godoc
// @Summary Add item to todos list
// @Description Add new item to list
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 404
// @Router /todo [post]
func (tc TodoController) Store(c echo.Context) error {
	var (
		item = new(ent.Todo)
		err  error
	)

	if err := c.Bind(&item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := todo.Store(item)
	helpers.Check(err)

	if tc.HttpType == "api" {
		return c.JSON(http.StatusOK, Response{Page: "Store", Payload: res})
	} else {
		return c.Redirect(http.StatusFound, "/web/todo")
	}

}

// Update godoc
// @Summary Update item in list TODOs
// @Description Add new item to list
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 404
// @Router /todo/{id} [post]
func (tc TodoController) Update(c echo.Context) error {
	var (
		id   = c.Param("id")
		item = new(ent.Todo)
		res  interface{}
		err  error
	)

	if err = c.Bind(item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	item.ID, err = strconv.Atoi(id)
	helpers.Check(err)
	res = todo.UpdateModel(item)

	if tc.HttpType == "api" {
		return c.JSON(http.StatusOK, Response{Page: "Update", Payload: res})
	} else {
		return c.Redirect(http.StatusFound, "/web/todo/"+c.Param("id"))
	}
}

// Delete godoc
// @Summary Delete item from TODOs list
// @Description Remove item from todos list
// @Tags todo
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 404
// @Router /todo/{id}/delete [get]
func (tc TodoController) Delete(c echo.Context) error {
	var (
		id, _ = strconv.Atoi(c.Param("id"))
	)
	todo.DelModel(id)

	if tc.HttpType == "api" {
		return c.JSON(http.StatusOK, Response{
			Page:    "Delete",
			Payload: "OK",
		})
	}
	return c.Redirect(http.StatusFound, "/web/todo/")
}
