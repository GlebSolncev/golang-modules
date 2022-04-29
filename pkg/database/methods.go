package database

import (
	"golang-modules/pkg/database/contracts"
	"golang-modules/pkg/database/sqlite"
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
