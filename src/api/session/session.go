package session

import (
	postgres "rest-server/src/db"
	"net/http"
	"encoding/json"
	"rest-server/src/middleware"
	"rest-server/src/services/crypto"
	"github.com/dgrijalva/jwt-go"
	"time"
	"log"
)

type Token struct {
	Token string `json:"token"`
}


func CreateSession(w http.ResponseWriter, r *http.Request)  {
	body := json.NewDecoder(r.Body)
	var parsedBody postgres.User
	err := body.Decode(&parsedBody)
	if err != nil {
		panic(err)
	}

	var user postgres.User
	err = postgres.Db.Model(&user).Where("email = ?", parsedBody.Email).Select()

	if err != nil {
		log.Println("Invalid user")
		http.Error(w, "Invalid", http.StatusUnauthorized)
		return
	}

	ok := crypto.ComparePasswords(user.Password, parsedBody.Password)
	if !ok {
		log.Println("Invalid password")
		http.Error(w, "Invalid", http.StatusUnauthorized)
		return
	}

	claims := middleware.TokenClaims{
		user.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(middleware.MySigningKey)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Token{ss})
}
