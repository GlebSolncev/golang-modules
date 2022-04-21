package main

import (
	"crud/internal/app"
	"crud/pkg/helpers"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	helpers.Check(err)
}

//go:generate echo "Build my app"
func main() {
	app.Start()
}
