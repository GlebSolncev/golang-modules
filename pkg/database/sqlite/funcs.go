package sqlite

import (
	"crud/pkg/database/contracts"
	"crud/pkg/ent"
	"crud/pkg/helpers"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
)

func (m *Method) Get() []byte {
	conn()
	data, _ := Client.Todo.Query().
		Select("id", "name", "slug", "status").
		All(ctx)

	closeConn()
	res, _ := json.Marshal(data)

	return res
}
func (m *Method) Save(data []byte, sType contracts.SaveType) bool {
	conn()
	var (
		item ent.Todo
		err  error
	)
	err = json.Unmarshal(data, &item)
	helpers.Check(err)

	switch sType {
	case contracts.New:
		addItem(item)
		break
	case contracts.Update:
		updItem(item)
		break
	case contracts.Delete:
		rmItem(item)
		break
	}

	closeConn()
	return true
}

func addItem(item ent.Todo) {
	todoCreate := Client.Todo.Create()

	if item.ID != 0 {
		todoCreate = todoCreate.SetID(item.ID)
	}
	if item.Name != "" {
		todoCreate = todoCreate.SetName(item.Name)
	}
	if item.Slug != "" {
		todoCreate = todoCreate.SetSlug(item.Slug)
	}

	todoCreate = todoCreate.SetStatus(item.Status)
	todoCreate.SaveX(ctx)

}

func updItem(item ent.Todo) {
	_, err := Client.Todo.
		UpdateOneID(item.ID).
		SetName(item.Name).
		SetSlug(item.Slug).
		SetStatus(item.Status).
		Save(ctx)

	helpers.Check(err)
}

func rmItem(item ent.Todo) {
	err := Client.Todo.
		DeleteOneID(item.ID).
		Exec(ctx)
	helpers.Check(err)
}
