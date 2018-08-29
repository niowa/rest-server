package db

import (
	"github.com/go-pg/pg"
	"os"
)

type User struct {
	Id string
	Email string
	Name string
	Password string
}

func ConnectToDb() *pg.DB {
	env := os.Getenv("ENV")
	var database string
	if (env == "DEV") {
		database = "go"
	} else {
		database = "go_test"
	}
	return pg.Connect(&pg.Options{
		User: "postgres",
		Password: "root",
		Database: database,
	})
}
