package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rumyantseva/cdays/internal/routing"
	"github.com/rumyantseva/cdays/internal/version"
)

func main() {
	log.Printf(
		"The application is starting, version is %s, build time is %s, commit is %v...",
		version.Release, version.BuildTime, version.Commit,
	)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("The port wasn't set")
	}

	diagPort := os.Getenv("DIAG_PORT")
	if diagPort == "" {
		log.Fatal("The diagnostics port wasn't set")
	}

	go func() {
		r := routing.NewBLRouter()
		log.Fatal(http.ListenAndServe(":"+port, r))
	}()

	{
		r := routing.NewDiagnosticsRouter()
		log.Fatal(http.ListenAndServe(":"+diagPort, r))
	}
}
