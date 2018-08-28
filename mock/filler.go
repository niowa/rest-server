package mock

import (
	postgres "rest-server/src/db"
	"github.com/satori/go.uuid"
	"rest-server/src/services/crypto"
	"rest-server/src/middleware"
	"github.com/dgrijalva/jwt-go"
	"time"
	"log"
)

type TestUser struct {
	User *postgres.User
	Token string
}

func FillDb() *TestUser {
	id, _ := uuid.NewV4()
	uuidAsString := id.String()
	name := "test"
	email := "test@test.com"
	password := "postman1"

	hashedPassword := crypto.HashAndSalt(password)

	newUser := TestUser{
		User: &postgres.User{
			Id: uuidAsString,
			Name: name,
			Email: email,
			Password: hashedPassword,
		},
	}

	db := postgres.ConnectToTestDb()
	defer db.Close()

	log.Printf("%+v", newUser.User)

	err := db.Insert(newUser.User)
	log.Printf("%+v", newUser)
	if err != nil {
		log.Println("Creation error")
		return nil
	}

	claims := middleware.TokenClaims{
		newUser.User.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(middleware.MySigningKey)

	newUser.Token = ss
	return &newUser
}
