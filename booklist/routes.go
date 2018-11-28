package booklist

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

func Main(port string) {

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	waitForShutdown := 15 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), waitForShutdown)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Shutting Down")
}
