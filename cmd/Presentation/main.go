package main

import (
    "log"
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Routers"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Database"
)

func main() {
    // Connexion à la base de données
    db, err := Database.ConnectDB()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    // Configuration des routes
    router := Routers.InitRoutes(db)

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}