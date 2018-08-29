package main

import (
	"log"
	"net/http"
	"rest-server/routes"
	"os"
)

type SelectUser struct {
	Id string
}

type TokenRequest struct {
	Token string
}

func main() {
	//ethereum.ConnectToEthereum()
	os.Setenv("ENV", "DEV")
	router := routes.Router()
	log.Fatal(http.ListenAndServe(":8000", router))
}
