package main

import (
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"golang-modules/internal/app"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	app.Start()
}
