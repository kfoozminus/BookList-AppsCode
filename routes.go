package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/", homePage).Methods("GET")
	Router.HandleFunc("/book", showBooks).Methods("GET")
	Router.HandleFunc("/book/{id}", showBook).Methods("GET")

	Router.HandleFunc("/book", authZ(addBook, true)).Methods("POST")
	Router.HandleFunc("/book/{id}", authZ(updateBook, true)).Methods("PUT")
	Router.HandleFunc("/book/{id}", authZ(deleteBook, true)).Methods("DELETE")
	Router.HandleFunc("/login", authZ(login, false)).Methods("POST")
	Router.HandleFunc("/logout", authZ(logout, true)).Methods("GET")
	Router.HandleFunc("/register", authZ(register, false)).Methods("POST")
}

func main() {

	err := http.ListenAndServe(":8080", Router)
	log.Fatal(err)
}
