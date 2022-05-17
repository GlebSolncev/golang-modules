package controllers

import (
	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger" // echo-swagger middleware
	_ "golang-modules/docs"            // swagger
	"golang-modules/internal/app/models"
	"golang-modules/pkg/ent"
	"golang-modules/pkg/helpers"
	"net/http"
	"strconv"
	"sync"
	_ "unsafe"
)

type (
	TodoController struct {
		Controllers
		Test     bool
		HttpType string
	}
)

var (
	todo = models.TodoModel{}
)

// Index godoc
// @Summary All todos items with statuses
// @Description Get todos list with status items
// @Tags Tоdo
// @Accept json
// @Produce json
// @Param   Authorization  header     string     true  "Token for auth"
// @Success 200 {object} Response
// @Failure 404
// @Router /api/todo [get]
func (tc TodoController) Index(c echo.Context) error {
	models.SetTypeWork(tc.Test)
	res, err := todo.GetAll()
	helpers.Check(err)

	if tc.HttpType == "api" {
		return c.JSON(http.StatusOK, res) //Response{NamePage: "Index", Payload: res})
	} else {
		return c.Render(http.StatusOK, "todo.tmpl", res)
	}
}

// Show godoc
// @Summary Item from todos list
// @Description Show item of Todos list
// @Tags Tоdo
// @Accept json
// @Produce json
// @Param	id				path	int		true 	"Tоdo ID"
// @Param   Authorization 	header	string	true  	"Token for auth"
// @Success 200 {object} Response
// @Failure 404
// @Router /api/todo/{id} [get]
func (tc TodoController) Show(c echo.Context) error {
	var (
		id  int
		err error
		res interface{}
	)

	models.SetTypeWork(tc.Test)
	id, err = strconv.Atoi(c.Param("id"))
	helpers.Check(err)
	res, err = todo.FindById(id)
	helpers.Check(err)

	if tc.HttpType == "api" {
		return c.JSON(http.StatusOK, Response{NamePage: "Show", Payload: res})
	} else {
		return c.Render(http.StatusOK, "todo.tmpl", res)
	}
}

// Store godoc
// @Summary Add item to todos list
// @Description Add new item to list
// @Tags Tоdo
// @Accept json
// @Produce json
// @Param   Authorization 	header	string		true  	"Token for auth"
// @Param	Body			body	string		true 	"Body for Tоdo item"
// @Success 200 {object} Response
// @Failure 404
// @Router /api/todo [post]
func (tc TodoController) Store(c echo.Context) error {
	var (
		item = new(ent.Todo)
	)

	if err := c.Bind(&item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(model *ent.Todo) {
		lock := sync.Mutex{}

		lock.Lock()
		item.ID, _ = todo.Store(model)
		lock.Unlock()

		wg.Done()
	}(item)

	wg.Wait()
	if tc.HttpType == "api" {
		return c.JSON(http.StatusCreated, Response{NamePage: "Store", Payload: item})
	} else {
		return c.Redirect(http.StatusCreated, "/web/todo")
	}

}

// Update godoc
// @Summary Update item in list TODOs
// @Description Add new item to list
// @Tags Todo
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 404
// @Router /api/todo/{id} [post]
func (tc TodoController) Update(c echo.Context) error {
	var (
		id   = c.Param("id")
		item = new(ent.Todo)
		res  interface{}
		err  error
	)

	models.SetTypeWork(tc.Test)
	if err = c.Bind(item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	item.ID, err = strconv.Atoi(id)
	helpers.Check(err)
	res = todo.UpdateModel(item)

	if tc.HttpType == "api" {
		return c.JSON(http.StatusOK, Response{NamePage: "Update", Payload: res})
	} else {
		return c.Redirect(http.StatusFound, "/web/todo/"+c.Param("id"))
	}
}

// Delete godoc
// @Summary Delete item from TODOs list
// @Description Remove item from todos list
// @Tags Todo
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 404
// @Router /api/todo/{id}/delete [get]
func (tc TodoController) Delete(c echo.Context) error {
	var (
		id, _ = strconv.Atoi(c.Param("id"))
	)

	models.SetTypeWork(tc.Test)
	todo.DelModel(id)

	if tc.HttpType == "api" {
		return c.JSON(http.StatusOK, Response{
			NamePage: "Delete",
			Payload:  "OK",
		})
	}
	return c.Redirect(http.StatusFound, "/web/todo/")
}
