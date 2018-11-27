package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Book struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

type Response struct {
	Success int    `json:"success"`
	Message string `json:"message,omitempty"`
	Book    []Book `json:"book,omitempty"`
}

type User struct {
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	Name          string `json:"name,omitempty"`
	LastSessionID string `json:"lastsessionid,omitempty"`
}

var userList = make(map[string]User)

var ind int

var bookList []Book
var mu sync.Mutex
var invalidResponse Response

func init() {
	ind++
	bookList = append(bookList, Book{ind, "Hundred Years of Solitude", "Gabriel García Márquez"})
	invalidResponse = Response{Success: 0, Message: "Invalid/Inefficient Information"}
}

func writeBad(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(invalidResponse)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the BookList RESTful API!")
}
