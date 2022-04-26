package main

import (
	"context"
	"crud/ent"
	"crud/internal/app"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	_ = godotenv.Load(".env")
}

func main() {

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app.Start()
}
