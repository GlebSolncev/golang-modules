package controllers

import (
	"github.com/labstack/echo/v4"
	"golang-modules/internal/app/models"
	"golang-modules/pkg/ent"
	"golang-modules/pkg/helpers"
	"net/http"
	"strconv"
)

type (
	StatusController struct {
		Controllers
	}
)

var (
	status = models.Status{}
)

// Index godoc
// @Summary All statuses
// @Description All statuses
// @Tags status
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /status [get]
func (StatusController) Index(c echo.Context) error {
	all, err := status.GetAll()
	helpers.Check(err)

	return c.JSON(http.StatusOK, Response{
		Page:    "Index",
		Payload: all,
	})
}

// Store godoc
// @Summary Add status
// @Description Add new item status
// @Tags status
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /status [post]
func (StatusController) Store(c echo.Context) error {
	var (
		item = new(ent.Status)
		err  error
	)
	if err := c.Bind(item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := status.Store(item)
	helpers.Check(err)

	return c.JSON(http.StatusOK, Response{Page: "Store", Payload: res})
}

// Show godoc
// @Summary Show item
// @Description Show item Status
// @Tags status
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /status/{id} [get]
func (StatusController) Show(c echo.Context) error {
	var (
		id  int
		err error
		res interface{}
	)
	id, err = strconv.Atoi(c.Param("id"))
	helpers.Check(err)
	res, err = status.FindById(id)
	helpers.Check(err)

	return c.JSON(http.StatusOK, Response{Page: "Show", Payload: res})
}

// Update godoc
// @Summary Update status
// @Description Update item Status
// @Tags status
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /status/{id} [put]
func (StatusController) Update(c echo.Context) error {
	var (
		id   = c.Param("id")
		item = new(ent.Status)
		res  interface{}
		err  error
	)

	if err = c.Bind(item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	item.ID, err = strconv.Atoi(id)
	helpers.Check(err)
	res = status.UpdateModel(item)

	return c.JSON(http.StatusOK, Response{Page: "Update", Payload: res})
}

// Delete godoc
// @Summary Remove status
// @Description Remove status
// @Tags status
// @Accept json
// @Produce json
// @Success 200
// @Router /status/{id} [delete]
func (StatusController) Delete(c echo.Context) error {
	var (
		id, _ = strconv.Atoi(c.Param("id"))
	)
	status.DelModel(id)

	return c.JSON(http.StatusOK, Response{
		Page:    "Delete",
		Payload: "OK",
	})
}
