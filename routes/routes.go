package routes

import (
	"github.com/gorilla/mux"
	"rest-server/src/middleware"
	"rest-server/src/api/profile"
	"rest-server/src/api/session"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware)
	router.HandleFunc("/profile", profile.GetProfile).Methods("GET")
	router.HandleFunc("/profile", profile.CreateProfile).Methods("POST")
	router.HandleFunc("/session", session.CreateSession).Methods("POST")

	return router
}
