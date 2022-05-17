package structirung

import "go/types"

type Struct struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name string
	Type types.Type
}
