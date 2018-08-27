package main

import (
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
	//ethereum.ConnectToEthereum()
	router := Router()
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware)
	router.HandleFunc("/profile", profile.GetProfile).Methods("GET")
	router.HandleFunc("/profile", profile.CreateProfile).Methods("POST")
	router.HandleFunc("/session", session.CreateSession).Methods("POST")

	return router
}
