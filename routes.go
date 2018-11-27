package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/book", showBooks).Methods("GET")
	r.HandleFunc("/book/{id}", showBook).Methods("GET")

	r.HandleFunc("/book", authZ(addBook, true)).Methods("POST")
	r.HandleFunc("/book/{id}", authZ(updateBook, true)).Methods("PUT")
	r.HandleFunc("/book/{id}", authZ(deleteBook, true)).Methods("DELETE")
	r.HandleFunc("/login", authZ(login, false)).Methods("GET")
	r.HandleFunc("/logout", authZ(logout, true)).Methods("GET")
	r.HandleFunc("/register", authZ(register, false)).Methods("POST")

	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
