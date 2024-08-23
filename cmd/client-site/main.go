package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    log.Println("Received request:", r.Method, r.URL.Path) // Ceci apparaîtra dans les logs
    fmt.Fprintf(w, "Hello, toi!")
}

func main() {
    log.Println("Starting server on :8081") // Ceci apparaîtra dans les logs
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8081", nil)) // Ceci apparaîtra dans les logs en cas d'erreur
}