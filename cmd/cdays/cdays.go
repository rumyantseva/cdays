package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("The application is starting...")

	http.HandleFunc("/", rootHandler())

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func rootHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	}
}
