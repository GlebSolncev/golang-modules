package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang-modules/internal/app/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	JSONbody = `{"name":"Work with Golang","slug":"work-with-golang","status":"InProgress"}`
)

func TestTodoController_Store(t *testing.T) {
	models.Init(".env.test")

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/todo", strings.NewReader(JSONbody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controllers := &TodoController{HttpType: "api"}

	if assert.NoError(t, controllers.Store(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestTodoController_Index(t *testing.T) {
	models.Init(".env.test")

	e := echo.New()
	controllers := TodoController{HttpType: "api"}
	req := httptest.NewRequest(http.MethodGet, "/api/todo", nil)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controllers.Index(c)

	var collec []map[string]interface{}
	var itemJSON map[string]interface{}

	if err := json.Unmarshal([]byte(rec.Body.String()), &collec); err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal([]byte(JSONbody), &itemJSON); err != nil {
		fmt.Println(err)
	}

	status := false
	for _, item := range collec {
		if itemJSON["name"] == item["name"] {
			status = true
			break
		}
	}
	if assert.True(t, status) {
		assert.Equal(t, http.StatusOK, rec.Code)

	}
}
