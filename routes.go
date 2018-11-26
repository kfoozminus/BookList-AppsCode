package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	r := mux.NewRouter()
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/book", showBooks).Methods("GET")
	r.HandleFunc("/book", addBook).Methods("POST")
	r.HandleFunc("/book/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/login", login).Methods("GET")
	r.HandlerFunc("/logout", logout).Methods("GET")
	r.HandlerFunc("/register", register).Methods("GET")

	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
