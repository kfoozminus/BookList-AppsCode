package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func showBook(w http.ResponseWriter, r *http.Request) {
	/*if isAuthorized(r) == false {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{Success: 0, Message: "No Authorization Provided"})
		return
	}*/
	mu.Lock()
	defer mu.Unlock()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		//not valid
		writeBad(w)
		return
	}

	for i, book := range bookList {

		if book.Id == id {

			json.NewEncoder(w).Encode(Response{Success: 1, Message: "Fetched Book Information Successfully", Book: []Book{bookList[i]}})
			return
		}
	}
	//not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{Success: 0, Message: "Book Not Found"})
}
