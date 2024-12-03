package Controllers

import (
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Errors"
)

type RefreshController struct{}

func NewRefreshController() *RefreshController {
    return &RefreshController{}
}

func (rc *RefreshController) HandleRefresh(w http.ResponseWriter, r *http.Request) {
    // Récupérer le refresh token depuis le cookie
    cookie, err := r.Cookie("refresh_token")
    if err != nil {
        if err == http.ErrNoCookie {
            http.Redirect(w, r, "/login?error=refresh_token_manquant", http.StatusSeeOther)
            return
        }
        http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
        return
    }

    refreshToken := cookie.Value
    if refreshToken == "" {
        http.Redirect(w, r, "/login?error=refresh_token_manquant", http.StatusSeeOther)
        return
    }

    // Valider le refresh token
    claims, err := Utils.ValidateJWT(refreshToken)
    if err != nil || claims["type"] != "refresh" {
        // Refresh token invalide ou expiré, supprimer le cookie et rediriger vers login
        Utils.SetRefreshTokenCookie(w, refreshToken)
        http.Redirect(w, r, "/login?error=invalid_refresh_token", http.StatusSeeOther)
        return
    }

    userIDFloat, ok := claims["user_id"].(float64)
    if !ok {
        http.Error(w, "Invalid token claims", http.StatusBadRequest)
        return
    }
    userID := int(userIDFloat)

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

    Utils.SetAccessTokenCookie(w, newAccessToken)
    Utils.SetRefreshTokenCookie(w, newRefreshToken)

    // Rediriger vers la page d'accueil ou une autre page protégée
    http.Redirect(w, r, "/", http.StatusSeeOther)
}