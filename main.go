package main

import (
	postgres "rest-server/src/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest-server/src/middleware"
	"rest-server/src/api/profile"
	"rest-server/src/api/session"
)

type SelectUser struct {
	Id string
}

type TokenRequest struct {
	Token string
}

func main() {
	postgres.ConnectToDb()
	//ethereum.ConnectToEthereum()
	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware)
	router.HandleFunc("/profile", profile.GetProfile).Methods("GET")
	router.HandleFunc("/profile", profile.CreateProfile).Methods("POST")
	router.HandleFunc("/session", session.CreateSession).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
