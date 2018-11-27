package main

import (
	"encoding/json"
	"net/http"
)

func addBook(w http.ResponseWriter, r *http.Request) {
	/*if isAuthorized(r) == false {

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{Success: 0, Message: "No Authorization Provided"})
		return
	}*/

	mu.Lock()
	defer mu.Unlock()

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err == nil {
		if book.Name == "" || book.Author == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Success: 0, Message: "Invalid/Inefficient information"})
			return
		}

		ind++
		book.Id = ind
		bookList = append(bookList, book)

		var _Book []Book
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(Response{Success: 1, Message: "Added Book Successfully!", Book: append(_Book, book)})

	} else {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Success: 0, Message: "Invalid/Inefficient information"})
	}
}
