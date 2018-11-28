package booklist

import (
	"encoding/json"
	"net/http"
)

func showBooks(w http.ResponseWriter, r *http.Request) {

	//if bookList is empty
	w.WriteHeader(http.StatusOK)
	if len(bookList) == 0 {
		json.NewEncoder(w).Encode(Response{Success: 1, Message: "No Book Added Yet"})
	} else {
		json.NewEncoder(w).Encode(Response{Success: 1, Message: "The Book List", Book: bookList})
	}
}
