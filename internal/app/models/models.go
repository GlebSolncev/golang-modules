package models

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
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
	c              *ent.Client
)

func Init(datasource string) {
	getDataSourceName(datasource)
	createDB()
}

func getDataSourceName(datasource string) {
	_ = godotenv.Load(path.GetBasePath(datasource))

	if os.Getenv("DB_DRIVER") == "" {
		panic("Env Not found")
	}

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
	err := c.Close()
	helpers.Check(err)
}

func createDB() {
	conn()
	defer closeConn()
	if err := c.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

func getTimeNow() time.Time {
	return time.Now()
}
