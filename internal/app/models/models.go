package models

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"golang-modules/pkg/ent"
	"golang-modules/pkg/helpers"
	"golang-modules/pkg/path"
	"os"
	"time"
)

type Model interface {
	GetAll() (interface{}, error)
	UpdateModel(model interface{}) interface{}
	DelModel(id int) bool
	FindById(id int) interface{}
	Store(model interface{}) (int, error)
}

var (
	dataSourceName string
	ctx            context.Context
	c              *ent.Client
)

func Init() {
	//_ = godotenv.Load(".env")
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

	c = cl
}

func closeConn() {
	defer c.Close()
}

func createDB() {
	conn()
	if err := c.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	closeConn()
}

func getTimeNow() time.Time {
	return time.Now()
}
