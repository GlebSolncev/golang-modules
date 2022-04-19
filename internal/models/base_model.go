package models

import "encoding/json"

func (Model) GetAll() []map[string]interface{} {
	var (
		todos []map[string]interface{}
		data  = manage.Get()
	)

	if len(data) >= 1 {
		err := json.Unmarshal(data, &todos)
		check(err)
	}

	return todos
}
