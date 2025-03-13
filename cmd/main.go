package main

import (
	"log"
	"nihal-innsof/portfolio/internal/server"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "[Portfolio-App] ", log.LstdFlags)

	port := 8080

	srvr, err := server.NewServer(logger, port)
	if err != nil {
		logger.Fatalf("Error creating server: %v", err)
		os.Exit(1)
	}
	if err := srvr.Start(); err != nil {
		logger.Fatalf("Error starting server: %v", err)
		os.Exit(1)
	}
}
