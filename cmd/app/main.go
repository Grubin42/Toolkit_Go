package main

import (
    "log"
    "net/http"
    "github.com/Grubin42/Toolkit_Go/internal/routers"
    "github.com/Grubin42/Toolkit_Go/pkg/database"
)

func main() {
    // Connexion à la base de données
    db, err := database.ConnectDB()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    // Configuration des routes
    router := routers.InitRoutes(db)

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}