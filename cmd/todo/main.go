package main

import (
	"github.com/GlebSolncev/golang-modules/internal/app"
	_ "github.com/GlebSolncev/golang-modules/internal/app/docs"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	_ = godotenv.Load(".env")
}

// @title Todo CRUD
// @version 1.0
// @description This is a simple CRUD for TODO list

// @contact.name GitHub
// @contact.url https://github.com/GlebSolncev/golang-modules

// @schemes http
// @host localhost:8080
// @BasePath /api
func main() {
	app.Start()
}
