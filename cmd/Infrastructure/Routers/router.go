package Routers

import (
    "database/sql"
    "net/http"
    "path/filepath"
    "log"
    "github.com/Grubin42/Toolkit_Go/cmd/Presentation/Controllers"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Middleware"
)

func InitRoutes(db *sql.DB) *http.ServeMux {
    router := http.NewServeMux()

    // Configuration des fichiers statiques
    assetsPath, err := filepath.Abs("cmd/Presentation/Assets")
    if err != nil {
        log.Fatal("Erreur lors de la résolution du chemin des assets :", err)
    }
    fs := http.FileServer(http.Dir(assetsPath))
    router.Handle("/assets/", http.StripPrefix("/assets/", fs))

    // Initialiser les contrôleurs
    homeController := Controllers.NewHomeController()
    adminController := Controllers.NewAdminController()
    registerController := Controllers.NewRegisterController(db)
    loginController := Controllers.NewLoginController(db)
    logoutController := Controllers.NewLogoutController()

    // Routes sans logs de débogage
    router.HandleFunc("/", homeController.HandleIndex)

    router.HandleFunc("/login", loginController.HandleIndex)

    router.HandleFunc("/register", registerController.HandleIndex)

    router.HandleFunc("/logout", logoutController.HandleLogout)
    
    router.Handle("/admin", Middleware.AuthMiddleware(http.HandlerFunc(adminController.HandleIndex)))

    return router
}