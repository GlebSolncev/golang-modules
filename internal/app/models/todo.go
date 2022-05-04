package models

import (
	"fmt"
	"github.com/GlebSolncev/golang-modules/pkg/ent"
	"github.com/GlebSolncev/golang-modules/pkg/ent/status"
	"github.com/GlebSolncev/golang-modules/pkg/ent/todo"
	"github.com/GlebSolncev/golang-modules/pkg/helpers"
)

type (
	Todo struct {
		Model
	}
)

func (Todo) GetStatuses() []ent.Status {
	return make([]ent.Status, 3)
	//return []todo.Status{
	//todo.StatusTodo, todo.StatusInProgress, todo.StatusDone, todo.StatusReview,
	//}
}

func (Todo) GetAll() (interface{}, error) {
	conn()
	data, err := c.Todo.
		Query().
		Select(todo.FieldID, todo.FieldSlug, todo.FieldName, status.FieldName).
		Select("id", "name", "slug").
		//WithStatus(func(q *ent.StatusQuery) {
		//	q.Select(status.FieldID, status.FieldName)
		//}).
		All(ctx)
	closeConn()

	fmt.Println(">>>>>>>> ", data)

	return data, err
}

func (Todo) UpdateModel(model interface{}) interface{} {
	m := model.(*ent.Todo)
	conn()
	model, err := c.Todo.
		UpdateOneID(m.ID).
		SetName(m.Name).
		SetSlug(m.Slug).
		//SetStatus(m.Status).
		Save(ctx)
	closeConn()
	helpers.Check(err)

	return model
}

func (Todo) DelModel(id int) bool {
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

func (tm Todo) FindById(id int) (interface{}, error) {
	conn()
	m, err := c.Todo.Query().Where(todo.ID(id)).First(ctx)
	closeConn()

	return m, err
}

func (Todo) Store(model interface{}) (interface{}, error) {
	conn()
	m := c.Todo.Create()
	for column, value := range model.(map[string]interface{}) {
		fmt.Println(column, value)
		switch column {
		case todo.FieldSlug:
			m = m.SetSlug(value.(string))
			break
		case todo.FieldName:
			m = m.SetName(value.(string))
			break
		}
	}
	res, err := m.Save(ctx)
	closeConn()

	return res, err
}
