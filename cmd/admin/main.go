// cmd/admin/main.go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Route pour la page d'administration
    http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello Admin")
    })

    // Démarrer le serveur sur le port 8080
    http.ListenAndServe(":8080", nil)
}