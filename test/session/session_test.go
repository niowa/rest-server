package session

import (
	postgres "rest-server/src/db"
	"testing"
	"log"
	"rest-server/mock"
	"encoding/json"
	"net/http"
	"bytes"
	"net/http/httptest"
	"rest-server/routes"
	"github.com/stretchr/testify/assert"
)

type tokenResponse struct {
	Token string
}

func TestSessionPost(t *testing.T) {
	ok := func() {
		log.Println("POST OK Test")
		userData := mock.FillDb()
		defer mock.CleanDb()

		user := postgres.User{
			Email: userData.User.Email,
			Name: userData.User.Name,
			Password: "postman1",
		}

		body, err := json.Marshal(user)
		request, _ := http.NewRequest("POST", "/session", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		request.Header.Set("Content-Type", "application/json")
		routes.Router().ServeHTTP(response, request)

		var parsedBody tokenResponse

		err = json.Unmarshal([]byte(response.Body.String()), &parsedBody)
		if err != nil {
			log.Println("body parser error")
			panic(err)
		}

		assert.IsType(t, parsedBody, tokenResponse{}, "Should return token")
	}
	ok()

	nonExistingUser := func() {
		log.Println("Non existing user Test")
		userData := mock.GenerateUser()
		defer mock.CleanDb()

		body, _ := json.Marshal(userData)
		request, _ := http.NewRequest("POST", "/session", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		request.Header.Set("Content-Type", "application/json")
		routes.Router().ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusUnauthorized, "Reject if user exists")
	}
	nonExistingUser()

	invalidParameters := func() {
		log.Println("Invalid Parameters Test")
		userData := new(postgres.User)
		body, _ := json.Marshal(userData)
		request, _ := http.NewRequest("POST", "/session", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		request.Header.Set("Content-Type", "application/json")
		routes.Router().ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusUnauthorized, "Reject if parameters are invalid")
	}
	invalidParameters()
}
