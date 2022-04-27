package database

import (
	"crud/pkg/database/contracts"
	"crud/pkg/database/memory"
)

type (
	NewMethod struct {
		Filename string
	}
)

var (
	defaultMethod contracts.DBMethods = &memory.Method{}
)

func (n NewMethod) Start() contracts.DBMethods {
	return defaultMethod
}
