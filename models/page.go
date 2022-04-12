package models

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"strconv"
)

type Page struct {
	Id   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type Pages struct {
	Pages []*Page `json:"page"`
}

var filename = "storage/page.json"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getCollect() []Page {
	var pages []Page
	var byteValue, _ = ioutil.ReadFile(filename)

	err := json.Unmarshal(byteValue, &pages)
	check(err)

	return pages
}

func GetPages() []Page {
	var pages = getCollect()

	return pages
}

func UpdatePage(c echo.Context) Page {
	var id, _ = strconv.Atoi(c.Param("id"))
	var page *Page = new(Page)
	err := c.Bind(page)
	check(err)
	var pages = getCollect()

	for k, item := range pages {
		//&pages
		if item.Id == id {
			//pages[k].Id = page.Id
			pages[k].Slug = page.Slug
			pages[k].Name = page.Name
		}
	}

	result, _ := json.Marshal(pages)
	err = ioutil.WriteFile(filename, result, 0644)
	return FindById(id)
}

func FindById(id int) Page {
	var resultPage Page
	var pages = getCollect()

	for _, page := range pages {
		if page.Id == id {
			resultPage = page
			break
		}
	}

	return resultPage
}

func StorePage(c echo.Context) bool {
	var pages []*Page
	var byteValue, _ = ioutil.ReadFile(filename)
	var page *Page = new(Page)

	err := c.Bind(page)
	check(err)
	err = json.Unmarshal(byteValue, &pages)
	check(err)
	pages = append(pages, page)
	result, _ := json.Marshal(pages)
	err = ioutil.WriteFile(filename, result, 0644)
	check(err)

	return true
}
