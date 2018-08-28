package main

import (
	"rest-server/mock"
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
)

//func TestProfile(t *testing.T) {
//	//req, err := http.Get("http://localhost:8000/profile")
//	//fmt.Printf("%+v\n", req.Status)
//	//fmt.Printf("%+v\n", err)
//
//	request, _ := http.NewRequest("GET", "/create", nil)
//	response := httptest.NewRecorder()
//	Router().ServeHTTP(response, request)
//	assert.Equal(t, 200, response.Code, "OK response is expected")
//}

func TestProfile(t *testing.T) {
	user := mock.FillDb()
	defer mock.CleanDb()
	request, _ := http.NewRequest("GET", "/profile", nil)
	response := httptest.NewRecorder()
	request.Header.Set("x-access-token", user.Token)
	Router().ServeHTTP(response, request)
	fmt.Println(response.Code)
}

