package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func updateBook(w http.ResponseWriter, r *http.Request) {
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
		writeBad(w)
		return
	}

	for i, book := range bookList {
		if book.Id == id {
			var updateBook Book
			err = json.NewDecoder(r.Body).Decode(&updateBook)

			if err == nil {

				bookList[i] = updateBook
				bookList[i].Id = id

				json.NewEncoder(w).Encode(Response{Success: 1, Message: "Updated Book Info Successfully!", Book: []Book{bookList[i]}})
			} else {
				writeBad(w)
			}
			return
		}
	}
	//not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{Success: 0, Message: "Book Not Found"})
}
