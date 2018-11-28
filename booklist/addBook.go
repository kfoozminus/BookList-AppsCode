package booklist

import (
	"encoding/json"
	"net/http"
)

func addBook(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	defer mu.Unlock()

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err == nil {
		if book.Name == "" || book.Author == "" {
			writeBad(w)
			return
		}

		ind++
		book.Id = ind
		bookList = append(bookList, book)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(Response{Success: 1, Message: "Added Book Successfully!", Book: []Book{book}})

	} else {
		writeBad(w)
	}
}
