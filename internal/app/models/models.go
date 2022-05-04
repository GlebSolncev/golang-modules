package models

import (
	"context"
	"fmt"
	"github.com/GlebSolncev/golang-modules/pkg/ent"
	"github.com/GlebSolncev/golang-modules/pkg/helpers"
	"github.com/GlebSolncev/golang-modules/pkg/path"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

var dataSourceName string
var ctx context.Context
var c *ent.Client

type Model interface {
	GetAll() []interface{}
	UpdateModel(model interface{}) interface{}
	DelModel(id int) bool
	FindById(id int) interface{}
	Store(model interface{}) bool
}

func init() {
	_ = godotenv.Load(".env")
	setDataSourceName()
	setContext()
	createDB()
}

func setContext() {
	ctx = context.Background()
}

func setDataSourceName() {
	host := os.Getenv("DB_HOST")

	if os.Getenv("DB_DRIVER") == "file" {
		host = path.GetBasePath(host)
	}

	dataSourceName = fmt.Sprintf("%s:%s%s%s?%s",
		os.Getenv("DB_DRIVER"),
		host,
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDITION"))
}

func conn() {
	cl, err := ent.Open("sqlite3", dataSourceName)
	helpers.Check(err)

	//defer helpers.Check(c.Close())

	c = cl
}

func closeConn() {
	//defer Client.Close()
}

func createDB() {
	c, err := ent.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer helpers.Check(c.Close())
	// Run the auto migration tool.
	if err := c.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}
