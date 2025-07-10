package main

import (
	"log"
	"os"

	"github.com/ElenaMask/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "morse-server: ", log.LstdFlags|log.Lshortfile)
	srv := server.NewServer(logger)
	if err := srv.HttpServer.ListenAndServe(); err != nil {
		logger.Fatalf("Server failed: %v", err)
	}
}
