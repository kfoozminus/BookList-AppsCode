package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {

	tests := []struct {
		verb         string
		url          string
		jsonBody     string
		expectedCode int
	}{
		{
			verb:         "GET",
			url:          "/",
			expectedCode: http.StatusOK,
		},
		{
			verb:         "GET",
			url:          "/book",
			expectedCode: http.StatusOK,
		},
		{
			verb:         "GET",
			url:          "/book/1",
			expectedCode: http.StatusOK,
		},
		{
			verb: "POST",
			url:  "/book",
			jsonBody: `{
							"name" : "To Kill a Mockingbird",
							"author" : "Harper Lee"
						}`,
			expectedCode: http.StatusUnauthorized,
		},
		{
			verb: "POST",
			url:  "/register",
			jsonBody: `{
							"username":"kfoozminus",
							"password":"kfoozminus",
							"name":"Jannatul Ferdows"
						}`,
			expectedCode: http.StatusOK,
		},
		{
			verb: "POST",
			url:  "/book",
			jsonBody: `{
							"name" : "To Kill a Mockingbird",
							"author" : "Harper Lee"
						}`,
			expectedCode: http.StatusUnauthorized,
		},
		{
			verb: "POST",
			url:  "/login",
			jsonBody: `{
						"username":"kfoozminus",
						"password":"kfoozminus"
						}`,
			expectedCode: http.StatusOK,
		},
		{
			verb: "POST",
			url:  "/book",
			jsonBody: `{
							"name" : "To Kill a Mockingbird",
							"author" : "Harper Lee"
						}`,
			expectedCode: http.StatusCreated,
		},
		{
			verb: "PUT",
			url:  "/book/1",
			jsonBody: `{
							"name" : "The Catcher in the Rye",
							"author" : "J. D. Salinger"
						}`,
			expectedCode: http.StatusOK,
		},
		{
			verb:         "DELETE",
			url:          "/book/1",
			expectedCode: http.StatusOK,
		},
		{
			verb:         "GET",
			url:          "/logout",
			expectedCode: http.StatusOK,
		},
		{
			verb: "POST",
			url:  "/book",
			jsonBody: `{
							"name" : "What I Talk About When I Talk About Running",
							"author" : "Haruki Murakami"
						}`,
			expectedCode: http.StatusUnauthorized,
		},
	}

	//var r *http.Request
	var lastSessionCookie http.Cookie
	for index, test := range tests {

		r, err := http.NewRequest(test.verb, test.url, strings.NewReader(test.jsonBody))
		if err != nil {
			log.Fatal(err)
		}
		r.AddCookie(&lastSessionCookie)

		w := httptest.NewRecorder()

		Router.ServeHTTP(w, r)

		for _, cookie := range w.Result().Cookies() {
			if cookie.Name == "SessionID" {
				lastSessionCookie = *cookie
			}
		}

		if w.Code != test.expectedCode {
			t.Errorf("Test %v : %v method %v didn't return %v\n", index, test.verb, test.url, test.expectedCode)
		}
	}
}
