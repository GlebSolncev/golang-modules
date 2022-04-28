package sqlite

import (
	"context"
	"crud/pkg/ent"
	"crud/pkg/helpers"
	"crud/pkg/path"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	_ = godotenv.Load(".env")
	setContext()
	setDataSourceName()
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

	DataSourceName = fmt.Sprintf("%s:%s%s%s?%s",
		os.Getenv("DB_DRIVER"),
		host,
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDITION"))
}

func conn() {
	c, err := ent.Open("sqlite3", DataSourceName)
	helpers.Check(err)

	Client = c
}

func closeConn() {
	//defer Client.Close()
}

func createDB() {
	client, err := ent.Open("sqlite3", DataSourceName)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}
