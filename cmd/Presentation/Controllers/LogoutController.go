// cmd/Presentation/Controllers/LogoutController.go
package Controllers

import (
    "net/http"
    "time"
    "database/sql"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Services"
    "log"
)

type LogoutController struct {
    refreshService *Services.RefreshService
}

// NewLogoutController initialise le LogoutController avec le RefreshService
func NewLogoutController(db *sql.DB) *LogoutController {
    return &LogoutController{
        refreshService: Services.NewRefreshService(db),
    }
}

// HandleLogout gère la déconnexion de l'utilisateur
func (lc *LogoutController) HandleLogout(w http.ResponseWriter, r *http.Request) {
    // Récupérer le refresh token depuis le cookie
    cookie, err := r.Cookie("refresh_token")
    if err == nil && cookie.Value != "" {
        // Révoquer le refresh token dans la base de données
        err = lc.refreshService.RevokeRefreshToken(cookie.Value)
        if err != nil {
            // Loguer l'erreur pour l'administration
            log.Printf("Erreur lors de la révocation du refresh token: %v\n", err)
            // Vous pouvez également afficher un message d'erreur générique à l'utilisateur si nécessaire
        }
    }

    // Supprimer les cookies de tokens en les définissant avec une date d'expiration passée
    http.SetCookie(w, &http.Cookie{
        Name:     "jwt_token",
        Value:    "",
        Expires:  time.Unix(0, 0),
        HttpOnly: true,
        Secure:   true, // À activer en production (HTTPS)
        SameSite: http.SameSiteStrictMode,
        Path:     "/",
    })

    http.SetCookie(w, &http.Cookie{
        Name:     "refresh_token",
        Value:    "",
        Expires:  time.Unix(0, 0),
        HttpOnly: true,
        Secure:   true, // À activer en production (HTTPS)
        SameSite: http.SameSiteStrictMode,
        Path:     "/refresh",
    })

    // Rediriger vers la page de login après la déconnexion
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}