package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
    // Get port from environment variable or default to 8080
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Define handlers
    http.HandleFunc("/", handleRoot)
    http.HandleFunc("/health", handleHealth)

    // Create a new server
    server := &http.Server{
        Addr:    ":" + port,
        Handler: nil, // Use the default ServeMux
    }

    // Start the server in a separate goroutine
    go func() {
        log.Printf("Server starting on port %s", port)
        if err := server.ListenAndServe(); err != nil {
            log.Fatalf("Could not listen on %s: %v", port, err)
        }
    }()

    // Set up a channel to listen for an interrupt or terminate signal from the OS
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM) // Combined signals

    <-stop // Wait for interrupt or terminate signal

    // Create a deadline to wait for. No need to name the context.
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel() // Call cancel as soon as shutdown is done.
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }

    log.Println("Server gracefully stopped")
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the Go Web App!")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "OK")
}
