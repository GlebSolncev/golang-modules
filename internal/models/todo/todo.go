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
	Todo Statuses = iota
	InProgress
	Done
	Review
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

func GetStatuses() []Statuses {
	ts := make([]Statuses, int(Review)+1)
	for i := 0; i <= int(Review); i++ {
		ts[i] = Statuses(i)
	}
	return ts
}

func GetAll() []Attributes {
	var (
		todos []Attributes
		data  = manage.Get()
	)

	if string(data) != "" {
		err := json.Unmarshal(data, &todos)
		helpers.Check(err)
		return todos
	}
	return make([]Attributes, 0)
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

			num, _ := strconv.ParseUint(todo.Status, 10, 32)
			todos[k].Status = Statuses(int(num)).String()
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

	num, _ := strconv.ParseUint(todo.Status, 10, 32)
	todo.Status = Statuses(int(num)).String()

	todos = append(todos, todo)
	data, err = json.Marshal(todos)
	helpers.Check(err)

	return manage.Save(data)
}
