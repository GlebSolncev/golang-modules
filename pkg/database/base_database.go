package database

import (
	"crud/pkg/path"
)

type (
	DBMethods interface {
		Get() []byte
		Save(data []byte) bool
	}

	NewMethod struct {
		Filename string
	}
)

var (
	defaultMethod DBMethods = &FileMethod{}
)

func (n NewMethod) Start() DBMethods {
	filename = path.GetBasePath(n.Filename)
	return defaultMethod
}

var filename = path.GetBasePath("storage/todo.json")

func check(err error) {
	if err != nil {
		panic(err)
	}
}
