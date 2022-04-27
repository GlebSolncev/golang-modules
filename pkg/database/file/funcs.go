package file

import (
	"crud/pkg/helpers"
	"encoding/json"
	"io/ioutil"
)

func init() {
	err := ioutil.WriteFile(filename, []byte{}, 0644)
	helpers.Check(err)
}

func (m *Method) Get() []byte {
	var byteValue, _ = ioutil.ReadFile(filename)

	return byteValue
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

func (m *Method) Save(data []byte) bool {
	helpers.Check(ioutil.WriteFile(filename, getCollectWithUpdIds(data), 0644))

	return true
}
