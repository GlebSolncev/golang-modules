package sqlite

import (
	"crud/ent"
	"crud/ent/todo"
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
func (m *Method) Save(data []byte) bool {
	conn()
	_, err := Client.Todo.Delete().Exec(ctx)
	helpers.Check(err)
	var items []map[string]interface{}
	err = json.Unmarshal(data, &items)
	helpers.Check(err)
	bulk := make([]*ent.TodoCreate, len(items))
	for i, val := range items {
		bulk[i] = Client.Todo.Create().
			SetName(val["name"].(string)).
			SetSlug(val["slug"].(string)).
			SetStatus(todo.Status(val["status"].(string)))
	}

	Client.Todo.CreateBulk(bulk...).SaveX(ctx)

	closeConn()
	return true
}
