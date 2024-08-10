// package main

// import (
//     "database/sql"
//     "time"
//     _ "github.com/lib/pq"
//     "log"
//     "github.com/gin-gonic/gin"
//     "toolkit_go/routes"
// )

// var db *sql.DB

// // connectToDB essaie de se connecter à la base de données PostgreSQL avec plusieurs tentatives
// func connectToDB() (*sql.DB, error) {
//     var db *sql.DB
//     var err error

//     // Essaie de se connecter 10 fois avec un délai de 2 secondes entre chaque essai
//     for i := 0; i < 10; i++ {
//         db, err = sql.Open("postgres", "postgres://toolkit_go:toolkit_go@postgres_db:5432/toolkit_go?sslmode=disable")
//         if err == nil {
//             err = db.Ping()
//             if err == nil {
//                 return db, nil
//             }
//         }
//         log.Println("Waiting for database to be ready...")
//         time.Sleep(2 * time.Second)
//     }
//     return nil, err
// }

// func main() {
//     var err error

//     // Appelle la fonction connectToDB pour se connecter à la base de données
//     db, err = connectToDB()
//     if err != nil {
//         log.Fatal("Failed to connect to the database:", err)
//     }
//     defer db.Close()

//     r := gin.Default()

//     // Charger les templates HTML depuis le répertoire correct
//     r.LoadHTMLGlob("templates/**/*.html")

//     // Charger les routes
//     routes.LoadRoutes(r)

//     // Démarrer le serveur sur le port 8080
//     r.Run(":8080")
// }

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