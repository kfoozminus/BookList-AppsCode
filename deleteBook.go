package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func deleteBook(w http.ResponseWriter, r *http.Request) {
	/*if isAuthorized(r) == false {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{Success: 0, Message: "No Authorization Provided"})
		return
	}*/

	mu.Lock()
	defer mu.Unlock()

	var delBook Book
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil {
		//not valid
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Success: 0, Message: "Invalid ID"})
		return
	}
	for i, book := range bookList {
		if book.Id == id {
			delBook = book
			bookList = append(bookList[:i], bookList[i+1:]...)

			var _Book []Book
			json.NewEncoder(w).Encode(Response{Success: 1, Message: "Deleted Book Successfully!", Book: append(_Book, delBook)})
			return
		}
	}
	//not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{Success: 0, Message: "Book Not Found"})
}
