package Routers

import (
    "database/sql"
 //   "Test/cmd/Grubin42/Toolkit_Go/Presentation/controllers"
    "net/http"
    "cmd/Presentation/controllers"
)

func InitRoutes(db *sql.DB) *http.ServeMux {
    router := http.NewServeMux()

    // Route pour les utilisateurs
    userController := controllers.NewUserController(db)

    // Routes
    router.HandleFunc("/", userController.HandleIndex)      // Utilise HandleIndex pour afficher le formulaire
    router.HandleFunc("/users", userController.HandleUsers) // Utilise HandleUsers pour gérer les utilisateurs
    router.HandleFunc("/users/list", userController.ListUsers) // Utilise listUsers pour lister les utilisateurs

    return router
}