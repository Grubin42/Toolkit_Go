package routers

import (
    "database/sql"
    "github.com/Grubin42/Toolkit_Go/internal/controllers"
    "net/http"
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