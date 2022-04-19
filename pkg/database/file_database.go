package database

import (
	"io/ioutil"
)

type (
	FileMethod struct {
		Method DBMethods
	}
)

func (m *FileMethod) Get() []byte {
	var byteValue, _ = ioutil.ReadFile(filename)

	return byteValue
}
func (m *FileMethod) Save(data []byte) bool {
	check(ioutil.WriteFile(filename, data, 0644))

	return true
}
