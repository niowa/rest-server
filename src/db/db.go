package db

import "github.com/go-pg/pg"

type User struct {
	Id string
	Email string
	Name string
	Password string
}

func ConnectToDb() *pg.DB {
	return pg.Connect(&pg.Options{
		User: "postgres",
		Password: "root",
		Database: "go",
	})
}

func ConnectToTestDb() *pg.DB {
	return pg.Connect(&pg.Options{
		User: "postgres",
		Password: "root",
		Database: "go_test",
	})
}
