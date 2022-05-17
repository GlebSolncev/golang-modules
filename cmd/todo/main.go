package main

import (
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/swaggo/echo-swagger" // echo-swagger middleware
	_ "golang-modules/docs"
	"golang-modules/internal/app"
	"golang-modules/internal/app/models"
)

// @title Tоdo CRUD
// @version 1.0
// @description This is a simple CRUD for Tоdo list

// @contact.name GitHub
// @contact.url https://github.com/GlebSolncev/golang-modules

// @schemes http
// @host localhost:8080
// @BasePath /
func main() {
	_ = godotenv.Load(".env")

	models.Init(false)
	app.Start()
}
