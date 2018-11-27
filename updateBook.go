package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func updateBook(w http.ResponseWriter, r *http.Request) {
	/*if isAuthorized(r) == false {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{Success: 0, Message: "No Authorization Provided"})
		return
	}*/
	mu.Lock()
	defer mu.Unlock()

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil {
		//not valid
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Success: 0, Message: "Invalid ID"})
		return
	}
	for i, book := range bookList {
		if book.Id == id {
			_ = json.NewDecoder(r.Body).Decode(&bookList[i])
			bookList[i].Id = id

			var _Book []Book
			json.NewEncoder(w).Encode(Response{Success: 1, Message: "Updated Book Info Successfully!", Book: append(_Book, bookList[i])})
			return
		}
	}
	//not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{Success: 0, Message: "Book Not Found"})
}
