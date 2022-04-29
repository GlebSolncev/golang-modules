package file

import (
	"golang-modules/pkg/database/contracts"
	"golang-modules/pkg/path"
)

type (
	Method struct {
		contracts.DBMethods
	}
)

var (
	filename = path.GetBasePath("storage/todo.json")
)
