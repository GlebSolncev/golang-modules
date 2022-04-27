package todo

import (
	"encoding/json"
	"testing"
)

func TestGetAll(t *testing.T) {
	all := GetAll()

	for _, item := range all {
		if item.Id == 0 {
			t.Log(item)
			t.Fatal("We have problem with json file. I have not an ID")
		}

		if item.Slug == "" {
			t.Log(item)
			t.Fatal("We have problem with json file. I have not a Slug")
		}

		if item.Name == "" {
			t.Log(item)
			t.Fatal("We have problem with json file. I have not a Name")
		}

		if item.Status == "" {
			t.Log(item)
			t.Fatal("We have problem with json file. I have not a Status")
		}
	}
}

func TestDeleteTodo(t *testing.T) {
	all := GetAll()
	collAftDel := DelModel(all[0].Id)

	if len(all) <= len(collAftDel) {
		t.Log("OLD data: ", all)
		t.Log("After delete item: ", collAftDel)
		t.Fatal("I cannot delete item from collect. Check pls")
	}
}

func TestStore(t *testing.T) {
	var (
		all     = GetAll()
		todo, _ = json.Marshal(all[len(all)-1])
		model   = new(Attributes)
	)

	_ = json.Unmarshal(todo, &model)
	for _, item := range all {
		item.Id = item.Id + 1
	}

	Store(model)

	if len(all) >= len(GetAll()) {
		t.Log("Old collect: ", all)
		t.Log("New collect: ", GetAll())
		t.Fatal("Problem with add new item")
	}
}

func TestStart(t *testing.T) {
	t.Run("GetAll", TestGetAll)
	t.Run("GetAll", TestDeleteTodo)
	t.Run("GetAll", TestStore)
}
