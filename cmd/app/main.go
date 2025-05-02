package main

import (
    "log"
    "net/http"
    "app/internal/handler"
)

func main() {
    mux := http.NewServeMux()

    // API routes
    mux.HandleFunc("/api/messages", handler.MessagesHandler)

    // Static file serving
    fs := http.FileServer(http.Dir("static"))
    mux.Handle("/", fs)

    log.Println("Listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
