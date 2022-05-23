package main

//Structuring

import (
	"golang-modules/internal/structuring"
	"path/filepath"
)

/**
Example:
	`go:generate main.go -name Human -all`
FLAGS:

-name [Name] 	- Name struct for find and work with it
-all 			- Get all fields and create getters/setters for it. Without this arg, I'm working with private fields.
*/
func getPath() string {
	dir, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	return dir
}

func main() {
	structuring.Execute(getPath())
}
