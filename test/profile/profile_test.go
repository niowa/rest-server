package profile

import (
	postgres "rest-server/src/db"
	"rest-server/routes"
	"rest-server/mock"
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"log"
	"bytes"
	"github.com/stretchr/testify/assert"
)

type tokenResponse struct {
	Token string
}

func TestProfileGet(t *testing.T) {
	ok := func() {
		log.Println("GET OK Test")
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
	ok()

	nonExistsUser := func() {
		log.Println("User does not exist Test")
		userData := mock.FillDb()
		mock.CleanDb()
		request, _ := http.NewRequest("GET", "/profile", nil)
		response := httptest.NewRecorder()
		request.Header.Set("x-access-token", userData.Token)
		routes.Router().ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusUnauthorized, "Reject if user is not exists")
	}
	nonExistsUser()

	notPassedToken := func() {
		log.Println("Token has not passed Test")
		request, _ := http.NewRequest("GET", "/profile", nil)
		response := httptest.NewRecorder()
		routes.Router().ServeHTTP(response, request)
		log.Println(response.Code)
		assert.Equal(t, response.Code, http.StatusUnauthorized, "Reject if token has not been passed")
	}
	notPassedToken()
}

func TestPorfilePost(t *testing.T) {
	ok := func() {
		log.Println("POST OK Test")
		userData := mock.GenerateUser()
		defer mock.CleanDb()

		body, err := json.Marshal(userData)
		request, _ := http.NewRequest("POST", "/profile", bytes.NewBuffer(body))
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

	duplicateUser := func() {
		log.Println("Duplicate User Test")
		userData := mock.FillDb()
		defer mock.CleanDb()

		userRequest := postgres.User{
			Email: userData.User.Email,
			Name: userData.User.Name,
			Password: "postman1",
		}

		body, _ := json.Marshal(userRequest)
		request, _ := http.NewRequest("POST", "/profile", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		request.Header.Set("Content-Type", "application/json")
		routes.Router().ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusUnprocessableEntity, "Reject if user is not exists")
	}
	duplicateUser()

	invalidParameters := func() {
		log.Println("Invalid Parameters Test")
		userData := new(postgres.User)
		body, _ := json.Marshal(userData)
		request, _ := http.NewRequest("POST", "/profile", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		request.Header.Set("Content-Type", "application/json")
		routes.Router().ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusBadRequest, "Reject if parameters are invalid")
	}
	invalidParameters()
}
