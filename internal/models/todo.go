package models

import (
	"crud/pkg/database"
	"encoding/json"
	"github.com/labstack/echo"
	"strconv"
)

var (
	manage = database.NewMethod{"storage/todo.json"}.Start()
	Parent = Model{}
)

func GetAllTodos() []Todo {
	var (
		todos []Todo
	)

	b, _ := json.Marshal(Parent.GetAll())
	err := json.Unmarshal(b, &todos)
	check(err)

	return todos
}

func UpdateTodo(c echo.Context) *Todo {
	var (
		id, _ = strconv.ParseFloat(c.Param("id"), 64)
		todos = Parent.GetAll()
		todo  = new(Todo)
		err   = c.Bind(todo)
	)
	check(err)

	for k, item := range todos {
		if id == item["id"] {
			//todos[k].Id = todo.Id
			todos[k]["slug"] = todo.Slug
			todos[k]["name"] = todo.Name
			todos[k]["status"] = todo.Status
		}
	}

	inrec, _ := json.Marshal(todos)
	manage.Save(inrec)

	return todo
}

func DelById(arr []Todo, id int) []Todo {
	var (
		newArr = make([]Todo, len(arr)-1)
		k      = 0
	)

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
		id, _ = strconv.Atoi(c.Param("id"))
		todos = GetAllTodos()
	)

	data, _ := json.Marshal(DelById(todos, id))
	manage.Save(data)
}

func FindById(id int) Todo {
	var (
		resultTodo Todo
		todos      = GetAllTodos()
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
