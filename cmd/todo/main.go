package main

import (
	"crud/internal/app"
)

//go:generate echo "Build my app"
func main() {
	app.Start()
}
