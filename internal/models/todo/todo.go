package todo

import (
	"crud/pkg/database"
	"crud/pkg/database/contracts"
	"crud/pkg/ent"
	"crud/pkg/ent/todo"
	"crud/pkg/helpers"
	"encoding/json"
)

var (
	manage = database.NewMethod{Filename: "storage/todo.json"}.Start()
)

func GetStatuses() []todo.Status {
	return []todo.Status{
		todo.StatusTodo, todo.StatusInProgress, todo.StatusDone, todo.StatusReview,
	}
}

func GetAll() []ent.Todo {
	var (
		todos []ent.Todo
		data  = manage.Get()
	)

	if string(data) != "" {
		err := json.Unmarshal(data, &todos)
		helpers.Check(err)

		return todos
	}
	return make([]ent.Todo, 0)
}

func SetModel(model *ent.Todo) *ent.Todo {

	inrec, _ := json.Marshal(model)
	manage.Save(inrec, contracts.Update)

	return model
}

func DelModel(id int) []ent.Todo {
	var (
		model ent.Todo
		err   error
		data  []byte
	)
	model.ID = id
	data, err = json.Marshal(model)
	helpers.Check(err)

	manage.Save(data, contracts.Delete)

	return []ent.Todo{}
}

func FindById(id int) ent.Todo {
	var (
		resultAttr ent.Todo
		ms         = GetAll()
	)

	for _, m := range ms {
		if m.ID == id {
			resultAttr = m
			break
		}
	}

	return resultAttr
}

func Store(model *ent.Todo) bool {
	dModel, _ := json.Marshal(model)

	return manage.Save(dModel, contracts.New)
}
