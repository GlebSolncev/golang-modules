package controllers

import (
	models "crud/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetPages(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name":  "Home",
		"pages": models.GetAll(),
		"page":  nil,
	})
}

func ShowPage(c echo.Context) error {
	var (
		id, _      = strconv.Atoi(c.Param("id"))
		statusCode = http.StatusOK
		page       = models.FindById(id)
	)

	if page.Slug == "" {
		statusCode = http.StatusNotFound
		page = models.Page{}
	}

	return c.Render(statusCode, "home.html", map[string]interface{}{
		"name":  "Home",
		"pages": models.GetAll(),
		"page":  page,
	})
}

func StorePage(c echo.Context) error {
	models.StorePage(c)

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name":  "Home",
		"pages": models.GetAll(),
		"page":  nil,
	})
}

func UpdatePage(c echo.Context) error {
	var (
		page       = models.UpdatePage(c)
		statusCode = http.StatusOK
	)

	if page.Slug == "" {
		statusCode = http.StatusNotFound
	}

	return c.Render(statusCode, "home.html", map[string]interface{}{
		"name":  "Home",
		"pages": models.GetAll(),
		"page":  page,
	})
}

func DeletePage(c echo.Context) error {
	var (
		statusCode = http.StatusOK
	)
	models.DeletePage(c)

	return c.Render(statusCode, "home.html", map[string]interface{}{
		"name":  "Home",
		"pages": models.GetAll(),
		"page":  nil,
	})
}
