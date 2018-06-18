package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	var blServer, diagServer http.Server

	srvErrs := make(chan error, 2)
	go func() {
		r := routing.NewBLRouter()
		blServer = http.Server{
			Addr:    ":" + port,
			Handler: r,
		}
		err := blServer.ListenAndServe()
		srvErrs <- err
	}()

	go func() {
		r := routing.NewDiagnosticsRouter()
		diagServer = http.Server{
			Addr:    ":" + diagPort,
			Handler: r,
		}
		err := diagServer.ListenAndServe()
		srvErrs <- err
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case killSignal := <-interrupt:
		log.Printf("Got %s. Stopping...", killSignal)
	case err := <-srvErrs:
		log.Printf("Got a server error: %s", err)
	}

	{
		ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
		blServer.Shutdown(ctx)
	}

	{
		ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
		diagServer.Shutdown(ctx)
	}
}
