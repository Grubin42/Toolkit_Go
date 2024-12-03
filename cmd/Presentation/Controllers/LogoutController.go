// cmd/Presentation/Controllers/LogoutController.go
package Controllers

import (
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
    "database/sql"
)

type LogoutController struct {}

// NewLogoutController initialise le LogoutController avec le RefreshService
func NewLogoutController(db *sql.DB) *LogoutController {
    return &LogoutController{}
}

// HandleLogout gère la déconnexion de l'utilisateur
func (lc *LogoutController) HandleLogout(w http.ResponseWriter, r *http.Request) {

    // Supprimer les cookies de tokens en les définissant avec une date d'expiration passée
    Utils.ClearTokenCookies(w)

    // Rediriger vers la page de login après la déconnexion
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}