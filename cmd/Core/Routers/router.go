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

    // Définir la route "/"
    router.HandleFunc("/", homeController.HandleIndex)

    return router
}