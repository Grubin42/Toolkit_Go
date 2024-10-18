package Routers

import (
    "database/sql"
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Presentation/Controllers"
)

func InitRoutes(db *sql.DB) *http.ServeMux {
    router := http.NewServeMux()

    // Initialiser le HomeController
    homeController := Controllers.NewHomeController()

    registerController := Controllers.NewRegisterController()
    loginController := Controllers.NewLoginController()

    // Définir la route "/"
    router.HandleFunc("/", homeController.HandleIndex)
    router.HandleFunc("/Register", registerController.HandleIndex)
    router.HandleFunc("/Login", loginController.HandleIndex)

    return router
}