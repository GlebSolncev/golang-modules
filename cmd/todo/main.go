package main

import (
	"crud/internal/app"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load(".env")
}

//go:generate echo "Build my app"
func main() {
	app.Start()
}
