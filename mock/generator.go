package mock

import (
	postgres "rest-server/src/db"
	"github.com/satori/go.uuid"
	"rest-server/src/services/crypto"
)

func GenerateUser() *postgres.User {
	id, _ := uuid.NewV4()
	uuidAsString := id.String()

	name := "test"
	email := "test@test.com"
	password := "postman1"

	hashedPassword := crypto.HashAndSalt(password)

	user := postgres.User{
		Id: uuidAsString,
		Name: name,
		Email: email,
		Password: hashedPassword,
	}
	return &user
}
