package controllers

import (
	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger" // echo-swagger middleware
	_ "golang-modules/docs"            // swagger
	"golang-modules/internal/app/models"
	"golang-modules/pkg/helpers"
	"net/http"
)

type (
	StatusController struct {
		Controllers
	}
)

var (
	status = models.StatusModel{}
)

// Index godoc
// @Summary All statuses
// @Description All statuses
// @Tags Status
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /api/status [get]
func (StatusController) Index(c echo.Context) error {
	all, err := status.GetAll()
	helpers.Check(err)

	return c.JSON(http.StatusOK, Response{
		NamePage: "Index",
		Payload:  all,
	})
}
