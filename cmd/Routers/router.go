package Routers

import (
    "Database/sql"
 //   "Test/cmd/Grubin42/Toolkit_Go/Presentation/controllers"
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Presentation/Controllers"
)

func InitRoutes(db *sql.DB) *http.ServeMux {
    router := http.NewServeMux()

    // Route pour les utilisateurs
    userController := Controllers.NewUserController(db)

    // Routes
    router.HandleFunc("/", userController.HandleIndex)      // Utilise HandleIndex pour afficher le formulaire
    router.HandleFunc("/users", userController.HandleUsers) // Utilise HandleUsers pour gérer les utilisateurs
    router.HandleFunc("/users/list", userController.ListUsers) // Utilise listUsers pour lister les utilisateurs

    return router
}