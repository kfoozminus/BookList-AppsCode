package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func main() {
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/book", showBooks).Methods("GET")
	router.HandleFunc("/book/{id}", showBook).Methods("GET")

	router.HandleFunc("/book", authZ(addBook, true)).Methods("POST")
	router.HandleFunc("/book/{id}", authZ(updateBook, true)).Methods("PUT")
	router.HandleFunc("/book/{id}", authZ(deleteBook, true)).Methods("DELETE")
	router.HandleFunc("/login", authZ(login, false)).Methods("POST")
	router.HandleFunc("/logout", authZ(logout, true)).Methods("GET")
	router.HandleFunc("/register", authZ(register, false)).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
