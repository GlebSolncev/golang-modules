package models

import (
	"github.com/GlebSolncev/golang-modules/pkg/ent"
	"github.com/GlebSolncev/golang-modules/pkg/ent/status"
	"github.com/GlebSolncev/golang-modules/pkg/helpers"
)

type (
	Status struct {
		Model
	}
)

func (Status) GetAll() (interface{}, error) {
	var (
		all []*ent.Status
		err error
	)
	conn()
	all, err = c.Status.Query().All(ctx)
	closeConn()

	return all, err
}

func (Status) Store(model interface{}) (interface{}, error) {
	var (
		m   *ent.Status
		err error
	)
	m, ok := model.(*ent.Status)
	if !ok {
		panic("I cannot convert interface to struct ent.Todo")
	}
	conn()
	m, err = c.Status.Create().SetName(m.Name).Save(ctx)
	closeConn()

	return m, err
}

func (Status) FindById(id int) (interface{}, error) {
	conn()
	m, err := c.Status.Query().Where(status.ID(id)).First(ctx)
	closeConn()

	return m, err
}

func (Status) UpdateModel(model interface{}) interface{} {
	m := model.(*ent.Status)
	conn()
	model, err := c.Status.
		UpdateOneID(m.ID).
		SetName(m.Name).
		Save(ctx)
	closeConn()
	helpers.Check(err)

	return model
}

func (Status) DelModel(id int) bool {
	var (
		err error
	)
	conn()
	err = c.Status.
		DeleteOneID(id).
		Exec(ctx)
	closeConn()
	helpers.Check(err)

	return true
}
