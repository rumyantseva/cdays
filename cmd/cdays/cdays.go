package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	go func() {
		r := routing.NewDiagnosticsRouter()
		log.Fatal(http.ListenAndServe(":"+diagPort, r))
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case killSignal := <-interrupt:
		log.Printf("Got %s. Stopping...", killSignal)
	}
}
