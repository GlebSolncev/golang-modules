package models

import (
	"fmt"
	"golang-modules/pkg/ent"
	"golang-modules/pkg/ent/todo"
	"golang-modules/pkg/helpers"
)

type (
	TodoModel struct {
		Model
	}
)

type TodoStatus int

const (
	Todo TodoStatus = iota
	InProgress
	Done
	Review
)

func (TodoModel) GetAll() (interface{}, error) {
	conn()
	data, err := c.Todo.
		Query().
		Select(todo.FieldID, todo.FieldSlug, todo.FieldName, todo.FieldCreatedAt, todo.FieldStatus).
		//Select("id", "name", "slug").
		//WithStatus(func(q *ent.StatusQuery) {
		//	q.Select(status.FieldID, status.FieldName)
		//}).
		All(ctx)
	closeConn()

	fmt.Println(">>>>>>>> ", data)

	return data, err
}

func (TodoModel) UpdateModel(model interface{}) interface{} {
	m := model.(*ent.Todo)
	conn()
	model, err := c.Todo.
		UpdateOneID(m.ID).
		SetName(*m.Name).
		SetSlug(m.Slug).
		//SetStatus(m.Status).
		Save(ctx)
	closeConn()
	helpers.Check(err)

	return model
}

func (TodoModel) DelModel(id int) bool {
	var (
		err error
	)
	conn()
	err = c.Todo.
		DeleteOneID(id).
		Exec(ctx)
	closeConn()
	helpers.Check(err)

	return true
}

func (tm TodoModel) FindById(id int) (interface{}, error) {
	conn()
	m, err := c.Todo.Query().Where(todo.ID(id)).First(ctx)
	closeConn()

	return m, err
}

func (TodoModel) Store(model *ent.Todo) (interface{}, error) {
	conn()
	m := c.Todo.Create()

	res, err := m.SetSlug(model.Slug).
		SetName(*model.Name).
		SetStatus(model.Status).
		SetCreatedAt(getTimeNow()).
		SetUpdatedAt(getTimeNow()).
		Save(ctx)

	closeConn()

	return res, err
}
