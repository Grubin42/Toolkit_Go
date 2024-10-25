package Controllers

import (
    "net/http"
	"time"
)

type LogoutController struct{}

func NewLogoutController() *LogoutController {
    return &LogoutController{}
}

func (lc *LogoutController) HandleLogout(w http.ResponseWriter, r *http.Request) {
    // Supprimer le cookie contenant le JWT en le définissant avec une date d'expiration passée
    http.SetCookie(w, &http.Cookie{
        Name:     "jwt_token",
        Value:    "",
        Expires:  time.Unix(0, 0), // Date d'expiration passée pour supprimer le cookie
        HttpOnly: true,
        Secure:   true, // Utiliser true si HTTPS est activé
    })

    // Rediriger vers la page de login après la déconnexion
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}