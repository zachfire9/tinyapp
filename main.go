package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func main() {
    // Function to log timestamp
    go func() {
        for {
            // Log the current timestamp
            log.Println("Current Timestamp:", time.Now().Format(time.RFC3339))
            time.Sleep(5 * time.Second) // Log every five seconds
        }
    }()

    // Index handler
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "Hello World!")
    })

    // Health check handler
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "OK")
    })

    // Start the HTTP server
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
