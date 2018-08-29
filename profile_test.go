package main

import (
	postgres "rest-server/src/db"
	"rest-server/routes"
	"rest-server/mock"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"log"
)

func TestGetProfile(t *testing.T) {
	userData := mock.FillDb()
	defer mock.CleanDb()
	request, _ := http.NewRequest("GET", "/profile", nil)
	response := httptest.NewRecorder()
	request.Header.Set("x-access-token", userData.Token)
	routes.Router().ServeHTTP(response, request)

	var parsedBody postgres.User

	err := json.Unmarshal([]byte(response.Body.String()), &parsedBody)
	if err != nil {
		log.Println("body parser")
		panic(err)
	}

	userForCompare := postgres.User{
		Id: userData.User.Id,
		Name: userData.User.Name,
		Email: userData.User.Email,
	}

	assert.Equal(t, userForCompare, parsedBody, "Should return user data")
}

