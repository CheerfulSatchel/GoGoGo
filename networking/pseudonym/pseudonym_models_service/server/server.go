package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func StartServer() {
	logger := log.New(os.Stdout, "pseudonym-database-api", log.LstdFlags)
	serverRoutes := AllRoutes()
	router := NewRouter(serverRoutes)

	server := &http.Server{
		Addr:         ":8081",
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		serverErr := server.ListenAndServe()
		if serverErr != nil {
			logger.Fatal(serverErr)
		}
	}()

	// For safely closing server...
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	receivedSignal := <-signalChannel
	logger.Printf("Received terminate signal %v, gracefully shutting down.", receivedSignal)

	shutdownContext, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	server.Shutdown(shutdownContext)
}
