package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rumyantseva/cdays/internal/routing"
)

func main() {
	log.Print("The application is starting...")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("The port wasn't set")
	}

	r := routing.NewBLRouter()
	log.Fatal(http.ListenAndServe(":"+port, r))
}
