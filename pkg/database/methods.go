package database

import (
	"crud/pkg/database/contracts"
	"crud/pkg/database/sqlite"
)

type (
	NewMethod struct {
		Filename string
	}
)

var (
	defaultMethod contracts.DBMethods = &sqlite.Method{}
)

func (n NewMethod) Start() contracts.DBMethods {
	return defaultMethod
}
