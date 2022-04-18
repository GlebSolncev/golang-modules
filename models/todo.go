package models

import (
	"crud/services/database"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"strconv"
)

type (
	Todo struct {
		Id     int    `json:"id"`
		Slug   string `json:"slug"`
		Name   string `json:"name"`
		Status string `json:"status"`
	}
	Todos struct {
		Todos []*Todo `json:"todo"`
	}
)

var manage = database.NewMethod{}.Start()

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetAll() []Todo {
	var (
		todos []Todo
		data  = manage.Get()
	)

	if len(data) >= 1 {
		err := json.Unmarshal(data, &todos)
		check(err)
	}

	return todos
}

func UpdateTodo(c echo.Context) *Todo {
	var (
		id, _ = strconv.Atoi(c.Param("id"))
		todos = GetAll()
		todo  = new(Todo)
		err   = c.Bind(todo)
	)
	check(err)

	for k, item := range todos {
		if id == item.Id {
			//todos[k].Id = todo.Id
			todos[k].Slug = todo.Slug
			todos[k].Name = todo.Name
			todos[k].Status = todo.Status
		}
	}

	inrec, _ := json.Marshal(todos)
	manage.Save(inrec)

	return todo
}

func RemoveFromList(arr []Todo, id int) []Todo {
	newArr := make([]Todo, len(arr)-1)
	k := 0
	for _, item := range arr {
		if item.Id != id {
			newArr[k] = item
			k++
		}
	}

	return newArr
}

func DeleteTodo(c echo.Context) {
	var (
		id, _    = strconv.Atoi(c.Param("id"))
		todos    = GetAll()
		newTodos []Todo
	)

	newTodos = RemoveFromList(todos, id)
	fmt.Println("X >>> ", newTodos)
	data, _ := json.Marshal(newTodos)
	manage.Save(data)
}

func FindById(id int) Todo {
	var (
		resultTodo Todo
		todos      = GetAll()
	)

	for _, todo := range todos {
		if todo.Id == id {
			resultTodo = todo
			break
		}
	}

	return resultTodo
}

func StoreTodo(c echo.Context) bool {
	var (
		todos []*Todo
		data  []byte = manage.Get()
		todo  *Todo  = new(Todo)
		err   error
	)

	if len(data) > 1 {
		err = json.Unmarshal(data, &todos)
		check(err)
	}

	err = c.Bind(todo)
	check(err)
	if len(todos) > 0 {
		todo.Id = todos[len(todos)-1].Id + 1
	} else {
		todo.Id = 1
	}
	todos = append(todos, todo)
	data, err = json.Marshal(todos)
	check(err)

	return manage.Save(data)
}
