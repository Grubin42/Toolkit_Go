package main

import (
    "log"
    "net/http"
    "path/filepath"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Routers"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Database"
)

func main() {
    // Connexion à la base de données
    db, err := database.ConnectDB()
    if err != nil {
        log.Fatal("Erreur lors de la connexion à la base de données :", err)
    }
    defer db.Close()

    // Initialisation des routes
    router := Routers.InitRoutes(db)

    // Servir les fichiers statiques
    fs := http.FileServer(http.Dir(filepath.Join("Presentation", "Assets")))
    router.Handle("/Assets/", http.StripPrefix("/Assets/", fs))

    log.Println("Serveur démarré sur le port :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Erreur du serveur : %v", err)
    }
}