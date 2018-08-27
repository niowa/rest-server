package profile

import (
	postgres "rest-server/src/db"
	"net/http"
	"github.com/gorilla/context"
	"encoding/json"
	"log"
	"github.com/satori/go.uuid"
	"rest-server/src/services/crypto"
	"rest-server/src/middleware"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserForResponse struct {
	Id interface{}
	Name interface{}
	Email interface{}
}

type Token struct {
	Token string `json:"token"`
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	user := UserForResponse{
		Id: context.Get(r, "id"),
		Name: context.Get(r, "name"),
		Email: context.Get(r, "email"),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	body := json.NewDecoder(r.Body)
	var parsedBody postgres.User
	err := body.Decode(&parsedBody)
	if err != nil {
		log.Println("body parser")
		panic(err)
	}

	var user postgres.User

	db := postgres.ConnectToDb()
	defer db.Close()

	err = db.Model(&user).Where("email = ?", parsedBody.Email).Select()

	if err == nil {
		log.Println("already had user")
		http.Error(w, "User with that email already exists", http.StatusUnprocessableEntity)
		return
	}

	id, _ := uuid.NewV4()
	uuidAsString := id.String()

	hashedPassword := crypto.HashAndSalt(parsedBody.Password)

	newUser := &postgres.User{
		Id: uuidAsString,
		Email: parsedBody.Email,
		Name: parsedBody.Name,
		Password: hashedPassword,
	}

	err = db.Insert(newUser)
	if err != nil {
		http.Error(w, "User was not added", http.StatusBadRequest)
		return
	}

	claims := middleware.TokenClaims{
		newUser.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(middleware.MySigningKey)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Token{ss})

}
