package main

import (
	"encoding/json"
	"net/http"
)

func showBooks(w http.ResponseWriter, r *http.Request) {

	/*if isAuthorized(r) == false {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{Success: 0, Message: "No Authorization Provided"})
		return
	}*/

	//if bookList is empty
	if len(bookList) == 0 {
		json.NewEncoder(w).Encode(Response{Success: 1, Message: "No Book Added Yet"})
	} else {
		json.NewEncoder(w).Encode(Response{Success: 1, Message: "The Book List", Book: bookList})
	}
}
