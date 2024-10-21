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
    adminController := Controllers.NewAdminController()

    registerController := Controllers.NewRegisterController(db)
    loginController := Controllers.NewLoginController()

    // Définir la route "/"
    router.HandleFunc("/", homeController.HandleIndex)
    router.HandleFunc("/admin", adminController.HandleIndex)
    
    router.HandleFunc("/register", registerController.HandleIndex)
    router.HandleFunc("/login", loginController.HandleIndex)

    return router
}