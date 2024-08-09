package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Crée une instance du routeur Gin
    r := gin.Default()

    // Définir une route pour la page d'accueil
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello, World!")
    })

    // Démarre le serveur sur le port 8080
    r.Run(":8080")
}