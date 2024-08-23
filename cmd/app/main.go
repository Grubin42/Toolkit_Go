package main

import (
    "net/http"
    "log"
)

func main() {
    http.HandleFunc("/", clientHandler)
    http.HandleFunc("/admin", adminHandler)
    
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func clientHandler(w http.ResponseWriter, r *http.Request) {
    // Gérer les requêtes du client
    w.Write([]byte("Hello, Client!"))
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
    // Gérer les requêtes de l'admin
    w.Write([]byte("Hello, Admin!"))
}