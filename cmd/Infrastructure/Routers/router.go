package Routers

import (
    "database/sql"
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Presentation/Controllers"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Middleware"
)

func InitRoutes(db *sql.DB) *http.ServeMux {
    router := http.NewServeMux()

    // Initialiser le HomeController
    homeController := Controllers.NewHomeController()
    adminController := Controllers.NewAdminController()
    registerController := Controllers.NewRegisterController(db)
    loginController := Controllers.NewLoginController(db)
    logoutController := Controllers.NewLogoutController()


    // Définir la route "/"
    router.HandleFunc("/", homeController.HandleIndex)

    router.HandleFunc("/register", registerController.HandleIndex)
    router.HandleFunc("/login", loginController.HandleIndex)
    router.HandleFunc("/logout", logoutController.HandleLogout)
    
    router.Handle("/admin", Middleware.AuthMiddleware(http.HandlerFunc(adminController.HandleIndex)))

    return router
}