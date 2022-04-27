package file

import (
	"crud/pkg/database/contracts"
	"crud/pkg/path"
)

type (
	Method struct {
		contracts.DBMethods
	}
)

var (
	filename = path.GetBasePath("storage/todo.json")
)
