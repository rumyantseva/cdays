package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Print("The application is starting...")

	r := mux.NewRouter()
	r.HandleFunc("/home", rootHandler())
	log.Fatal(http.ListenAndServe(":8000", r))
}

func rootHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	}
}
