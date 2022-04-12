package controllers

import (
	"crud/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetPages(c echo.Context) error {
	return c.JSON(http.StatusOK, models.GetPages())
}

func ShowPage(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))
	var statusCode = http.StatusOK
	var page = models.FindById(id)

	if page.Slug == "" {
		statusCode = http.StatusNotFound
		page = models.Page{}
	}

	return c.JSON(statusCode, page)
}

func StorePage(c echo.Context) error {
	models.StorePage(c)

	return c.JSON(http.StatusOK, models.GetPages())
}

func UpdatePage(c echo.Context) error {
	var page = models.UpdatePage(c)
	var statusCode = http.StatusOK
	if page.Slug == "" {
		statusCode = http.StatusNotFound
		page = models.Page{}
	}

	return c.JSON(statusCode, page)
}
