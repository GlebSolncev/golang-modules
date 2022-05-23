package models

import (
	"context"
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

func (TodoModel) GetAll() ([]*ent.Todo, error) {
	conn()
	defer closeConn()
	data, err := c.Todo.
		Query().
		Select(todo.FieldID, todo.FieldSlug, todo.FieldName, todo.FieldCreatedAt, todo.FieldStatus).
		All(context.Background())

	return data, err
}

func (TodoModel) UpdateModel(model interface{}) interface{} {
	m := model.(*ent.Todo)
	conn()
	defer closeConn()
	model, err := c.Todo.
		UpdateOneID(m.ID).
		SetName(m.Name).
		SetSlug(m.Slug).
		SetStatus(m.Status).
		Save(context.Background())

	helpers.Check(err)

	return model
}

func (TodoModel) DelModel(id int) bool {
	var (
		err error
	)
	conn()
	defer closeConn()
	err = c.Todo.
		DeleteOneID(id).
		Exec(context.Background())
	helpers.Check(err)

	return true
}

func (tm TodoModel) FindById(id int) (interface{}, error) {
	conn()
	defer closeConn()
	m, err := c.Todo.Query().Where(todo.ID(id)).First(context.Background())

	return m, err
}

func (TodoModel) Store(model *ent.Todo) (int, error) {
	conn()
	defer c.Close()
	m := c.Todo.Create()

	fillable := m.SetSlug(model.Slug).
		SetName(model.Name).
		SetStatus(model.Status).
		SetCreatedAt(getTimeNow()).
		SetUpdatedAt(getTimeNow())

	if model.ID > 0 {
		fillable = fillable.SetID(model.ID)
	}

	res, err := fillable.Save(context.Background())

	if err != nil {
		return 0, err
	}

	return res.ID, err
}
