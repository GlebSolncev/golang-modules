package main

import (
	"crud/internal/app"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	app.Start()
}
