package mock

import (
	postgres "rest-server/src/db"
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
	user := GenerateUser()

	userData := TestUser{
		User: user,
	}
	db := postgres.ConnectToDb()
	defer db.Close()

	err := db.Insert(userData.User)

	if err != nil {
		log.Println("Creation error")
		return nil
	}

	claims := middleware.TokenClaims{
		userData.User.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(middleware.MySigningKey)
	userData.Token = tokenString

	return &userData
}
