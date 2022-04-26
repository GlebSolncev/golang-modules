package todo

import (
	"crud/pkg/database"
	"crud/pkg/helpers"
	"encoding/json"
	"github.com/labstack/echo"
	"strconv"
)

//go:generate stringer --type=Statuses

type Statuses int

const (
	Draft Statuses = iota
	Start
	InProcess
	Review
	Done
)

type (
	Attributes struct {
		Id     int    `json:"id"`
		Slug   string `json:"slug"`
		Name   string `json:"name"`
		Status string `json:"status" binding:"required,gte=0,lte=3"`
	}
)

var (
	manage = database.NewMethod{Filename: "storage/todo.json"}.Start()
)

func GetAll() []Attributes {
	var (
		todos []Attributes
		data  = manage.Get()
	)

	err := json.Unmarshal(data, &todos)
	helpers.Check(err)

	return todos
}

func SetModel(c echo.Context) *Attributes {
	var (
		id, _ = strconv.Atoi(c.Param("id"))
		data  = manage.Get()
		todos []Attributes
		todo  = new(Attributes)
		err   error
	)

	err = c.Bind(todo)
	helpers.Check(err)

	err = json.Unmarshal(data, &todos)
	helpers.Check(err)

	for k, item := range todos {
		if id == item.Id {
			todos[k].Slug = todo.Slug
			todos[k].Name = todo.Name
			todos[k].Status = todo.Status
		}
	}

	inrec, _ := json.Marshal(todos)
	manage.Save(inrec)

	return todo
}

func DelModel(id int) []Attributes {
	var (
		todos  = GetAll()
		newArr = make([]Attributes, len(todos)-1)
		k      = 0
	)

	for _, item := range todos {
		if item.Id != id {
			newArr[k] = item
			k++
		}
	}

	data, _ := json.Marshal(newArr)
	manage.Save(data)

	return newArr
}

func FindById(id int) Attributes {
	var (
		resultAttr Attributes
		todos      = GetAll()
	)

	for _, todo := range todos {
		if todo.Id == id {
			resultAttr = todo
			break
		}
	}

	return resultAttr
}

func Store(todo *Attributes) bool {
	var (
		todos []*Attributes
		data  = manage.Get()
		err   error
	)

	if len(data) > 1 {
		err = json.Unmarshal(data, &todos)
		helpers.Check(err)
	}

	if len(todos) > 0 {
		todo.Id = todos[len(todos)-1].Id + 1
	} else {
		todo.Id = 1
	}

	todos = append(todos, todo)
	data, err = json.Marshal(todos)
	helpers.Check(err)

	return manage.Save(data)
}
