package main

import (
	"log"
	"net/http"

	"github.com/rumyantseva/cdays/internal/routing"
)

func main() {
	log.Print("The application is starting...")

	r := routing.NewBLRouter()
	log.Fatal(http.ListenAndServe(":8000", r))
}
