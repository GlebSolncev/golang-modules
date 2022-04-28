package memory

import (
	"crud/pkg/database/contracts"
	"crud/pkg/helpers"
	"encoding/json"
)

func (m *Method) Get() []byte {
	return db.Data
}

func getCollectWithUpdIds(data []byte) []byte {
	var items []map[string]interface{}
	err := json.Unmarshal(data, &items)
	helpers.Check(err)
	for i := range items {
		items[i]["id"] = i + 1
	}
	data, err = json.Marshal(items)
	helpers.Check(err)

	return data
}

func (m *Method) Save(data []byte, sType contracts.SaveType) bool {
	db.Data = getCollectWithUpdIds(data)

	return true
}
