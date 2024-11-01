// cmd/Presentation/Controllers/RefreshController.go
package Controllers

import (
    "net/http"
	"database/sql"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Services"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Errors"
    "time"
)

type RefreshController struct {
    refreshService *Services.RefreshService
}

func NewRefreshController(db *sql.DB) *RefreshController {
    return &RefreshController{
        refreshService: Services.NewRefreshService(db),
    }
}

func (rc *RefreshController) HandleRefresh(w http.ResponseWriter, r *http.Request) {
    // Récupérer le refresh token depuis le cookie
    cookie, err := r.Cookie("refresh_token")
    if err != nil {
        if err == http.ErrNoCookie {
            // Aucun refresh token trouvé, rediriger vers la page de login
            http.Redirect(w, r, "/login?error=refresh_token_manquant", http.StatusSeeOther)
            return
        }
        // Autre erreur, afficher une erreur interne du serveur
        http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
        return
    }

    refreshToken := cookie.Value
    if refreshToken == "" {
        http.Redirect(w, r, "/login?error=refresh_token_manquant", http.StatusSeeOther)
        return
    }

    // Valider le refresh token
    userID, err := rc.refreshService.ValidateRefreshToken(refreshToken)
    if err != nil {
        // Si le refresh token est invalide ou expiré, supprimer le cookie et rediriger vers login
        http.SetCookie(w, &http.Cookie{
            Name:     "refresh_token",
            Value:    "",
            Expires:  time.Unix(0, 0),
            HttpOnly: true,
            Secure:   true,
            SameSite: http.SameSiteStrictMode,
            Path:     "/refresh",
        })
        http.Redirect(w, r, "/login?error=invalid_refresh_token", http.StatusSeeOther)
        return
    }

    // Générer un nouveau access token
    newAccessToken, err := Utils.GenerateAccessToken(userID)
    if err != nil {
        http.Error(w, Errors.ErrorTokenGeneration, http.StatusInternalServerError)
        return
    }

    // Générer un nouveau refresh token pour la rotation
    newRefreshToken, err := Utils.GenerateRefreshToken(userID)
    if err != nil {
        http.Error(w, Errors.ErrorRefreshTokenGeneration, http.StatusInternalServerError)
        return
    }

    // Enregistrer le nouveau refresh token dans la base de données
    expiresAt := time.Now().Add(Utils.GetRefreshTokenExpiration())
    err = rc.refreshService.SaveRefreshToken(userID, newRefreshToken, expiresAt)
    if err != nil {
        http.Error(w, Errors.ErrorRefreshTokenSave, http.StatusInternalServerError)
        return
    }

    // Révoquer l'ancien refresh token
    err = rc.refreshService.RevokeRefreshToken(refreshToken)
    if err != nil {
        http.Error(w, Errors.ErrorRevokeAllTokens, http.StatusInternalServerError)
        return
    }

    // Mettre à jour les cookies avec les nouveaux tokens
    http.SetCookie(w, &http.Cookie{
        Name:     "jwt_token",
        Value:    newAccessToken,
        Expires:  time.Now().Add(Utils.GetAccessTokenExpiration()),
        HttpOnly: true,
        Secure:   false, // À activer en production
        SameSite: http.SameSiteStrictMode,
        Path:     "/",
    })

    http.SetCookie(w, &http.Cookie{
        Name:     "refresh_token",
        Value:    newRefreshToken,
        Expires:  expiresAt,
        HttpOnly: true,
        Secure:   false, // À activer en production
        SameSite: http.SameSiteStrictMode,
        Path:     "/refresh",
    })

    // Rediriger vers la page d'accueil ou une autre page protégée
    http.Redirect(w, r, "/", http.StatusSeeOther)
}