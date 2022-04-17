package models

import (
	"crud/services/database"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"strconv"
)

type (
	Page struct {
		Id   int    `json:"id"`
		Slug string `json:"slug"`
		Name string `json:"name"`
	}
	Pages struct {
		Pages []*Page `json:"page"`
	}
)

var manage = database.NewMethod{}.Start()

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetAll() []Page {
	var (
		pages []Page
		data  = manage.Get()
	)

	if len(data) >= 1 {
		err := json.Unmarshal(data, &pages)
		check(err)
	}

	return pages
}

func UpdatePage(c echo.Context) *Page {
	var (
		id, _ = strconv.Atoi(c.Param("id"))
		pages = GetAll()
		page  = new(Page)
		err   = c.Bind(page)
	)
	check(err)

	for k, item := range pages {
		if id == item.Id {
			//pages[k].Id = page.Id
			pages[k].Slug = page.Slug
			pages[k].Name = page.Name
		}
	}

	inrec, _ := json.Marshal(pages)
	manage.Save(inrec)

	return page
}

func RemoveFromList(arr []Page, id int) []Page {
	new_arr := make([]Page, len(arr)-1)
	k := 0
	for _, item := range arr {
		if item.Id != id {
			new_arr[k] = item
			k++
		}
	}

	return new_arr
}

func DeletePage(c echo.Context) {
	var (
		id, _    = strconv.Atoi(c.Param("id"))
		pages    = GetAll()
		newPages []Page
	)

	newPages = RemoveFromList(pages, id)
	fmt.Println("X >>> ", newPages)
	data, _ := json.Marshal(newPages)
	manage.Save(data)
}

func FindById(id int) Page {
	var (
		resultPage Page
		pages      = GetAll()
	)

	for _, page := range pages {
		if page.Id == id {
			resultPage = page
			break
		}
	}

	return resultPage
}

func StorePage(c echo.Context) bool {
	var (
		pages []*Page
		data  []byte = manage.Get()
		page  *Page  = new(Page)
		err   error
	)

	if len(data) > 1 {
		err = json.Unmarshal(data, &pages)
		check(err)
	}

	err = c.Bind(page)
	check(err)
	if len(pages) > 0 {
		page.Id = pages[len(pages)-1].Id + 1
	} else {
		page.Id = 1
	}
	pages = append(pages, page)
	data, err = json.Marshal(pages)
	check(err)

	return manage.Save(data)
}
