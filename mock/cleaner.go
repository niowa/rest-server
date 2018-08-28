package mock

import (
	postgres "rest-server/src/db"
	"log"
)

func CleanDb() {
	db := postgres.ConnectToTestDb()
	defer db.Close()

	sqlStatment := `
	DELETE from users
	WHERE name = 'test';`

	_, err := db.Exec(sqlStatment)

	if err != nil {
		log.Println("Removing error")
		return
	}
}
