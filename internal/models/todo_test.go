package models

import (
	"testing"
)

func TestGetAllTodos(t *testing.T) {
	all := Model{}.GetAll()

	for _, item := range all {
		if item["id"] == nil {
			t.Log(item)
			t.Fatal("We have problem with json file. I have not an ID")
		}

		if item["slug"] == nil {
			t.Log(item)
			t.Fatal("We have problem with json file. I have not a Slug")
		}

		if item["name"] == nil {
			t.Log(item)
			t.Fatal("We have problem with json file. I have not a Name")
		}

		if item["status"] == nil {
			t.Log(item)
			t.Fatal("We have problem with json file. I have not a Status")
		}
	}
}

func TestDeleteTodo(t *testing.T) {
	all := GetAllTodos()
	id := all[0].Id
	aDel := DelById(all, id)

	if len(all) <= len(aDel) {
		t.Log("OLD data: ", all)
		t.Log("After delete item: ", aDel)
		t.Fatal("I cannot delete item from collect. Check pls")
	}
}
