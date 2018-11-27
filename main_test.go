package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
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
		/*{
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
			expectedCode: http.StatusOK,
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
		},*/
	}

	for index, test := range tests {
		r, err := http.NewRequest(test.verb, test.url, strings.NewReader(test.jsonBody))
		//spew.Dump(r)
		if err != nil {
			log.Fatal(err)
		}
		w := httptest.NewRecorder()

		router.ServeHTTP(w, r)
		spew.Dump(w)
		//fmt.Printf("%v %v %v\n", index, w.Code, test.expectedCode)

		if w.Code != test.expectedCode {
			t.Errorf("Test %v : %v didn't return %v\n", index, test.url, test.expectedCode)
			//t.Errorf("%v didn't return %v for %v", test.url, test.expectedCode, test.jsonBody)
		}
	}
}
