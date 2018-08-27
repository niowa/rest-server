package middleware

import (
	"net/http"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	postgres "rest-server/src/db"
)

var MySigningKey = []byte("f4vb8fJu9hE9XfX6szY5awQU")

type TokenClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

var allowedPostRoutes = map [string]string {
	"/profile": "",
	"/session": "",
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := allowedPostRoutes[r.URL.String()]; ok && r.Method == "POST" {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := r.Header.Get("x-access-token")

		token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(MySigningKey), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println(err)
			return
		}

		if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
			db := postgres.ConnectToDb()
			defer db.Close()

			var user postgres.User
			err = db.Model(&user).Where("id = ?", claims.Id).Select()
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Println(err)
				return
			}
			context.Set(r, "email", user.Email)
			context.Set(r, "name", user.Name)
			context.Set(r, "password", user.Password)
			context.Set(r, "id", user.Id)
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}
