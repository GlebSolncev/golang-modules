package models

import (
	"github.com/stretchr/testify/assert"
	"golang-modules/pkg/ent"
	"golang-modules/pkg/ent/todo"
	"golang-modules/pkg/helpers"
	"testing"
)

var todoModel = TodoModel{}

func upTestDatabase() {
	Init(".env.test")
}

func TestTodoModel_Store(t *testing.T) {
	var testCases = []struct {
		name       string // Name test case
		wantReturn int    // equal
		item       *ent.Todo
		wantErr    bool
	}{
		{
			name: "New item",
			item: &ent.Todo{
				ID:     10000,
				Name:   "Item 1",
				Slug:   "test description",
				Status: todo.StatusInProgress,
			},
			wantReturn: 10000,
		},
		{
			name: "Check unique ID. Add new item with old ID",
			item: &ent.Todo{
				ID:     10000,
				Name:   "Err item",
				Slug:   "Error...",
				Status: todo.StatusInProgress,
			},
			wantErr: true,
		},
		{
			name: "Already exist",
			item: &ent.Todo{
				ID:     10001,
				Name:   "Item 2",
				Slug:   "item-2",
				Status: todo.StatusTodo,
			},
			wantReturn: 10001,
		},
	}
	upTestDatabase()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var (
				result int
				err    error
			)

			result, err = todoModel.Store(testCase.item)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.wantReturn, result)
			}

		})
	}
}

func TestTodoModel_DelModel(t *testing.T) {
	upTestDatabase()
	all, err := todoModel.GetAll()
	helpers.Check(err)
	for _, model := range all {
		t.Run(model.Name, func(t *testing.T) {
			todoModel.DelModel(model.ID)
		})
	}
}
