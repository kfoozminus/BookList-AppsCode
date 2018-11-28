package booklist

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func deleteBook(w http.ResponseWriter, r *http.Request) {

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

			var delBook Book
			delBook = book
			bookList = append(bookList[:i], bookList[i+1:]...)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Response{Success: 1, Message: "Deleted Book Successfully!", Book: []Book{delBook}})
			return
		}
	}
	//not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{Success: 0, Message: "Book Not Found"})
}
