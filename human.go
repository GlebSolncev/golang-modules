package golang_modules

import "golang-modules/pkg/path"

//go:generate go run ./cmd/structuring/main.go -name Human

type (
	Human struct {
		FirstName string
		LastName  string
		Age       int
		address   path.HelloWorld
		addr      Address
		passport  string
	}

	Address struct {
		Home   string
		Street string
	}
)
