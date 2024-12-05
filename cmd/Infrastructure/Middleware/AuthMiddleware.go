package Middleware

import (
    "net/http"
    "errors"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Errors"
)

func AuthMiddleware() func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Récupérer le access token depuis le cookie
            cookie, err := r.Cookie("jwt_token")
            if err != nil {
                if errors.Is(err, http.ErrNoCookie) {
                    http.Redirect(w, r, "/login", http.StatusSeeOther)
                    return
                }
                http.Error(w, Errors.ErrInternalServerError.Error(), http.StatusInternalServerError)
                return
            }

            accessToken := cookie.Value
            if accessToken == "" {
                http.Redirect(w, r, "/login", http.StatusSeeOther)
                return
            }

            // Valider le access token
            claims, err := Utils.ValidateJWT(accessToken)
            if err != nil {
                // Si le token est expiré, tenter de rafraîchir
                if err.Error() == "token invalide ou expiré. Veuillez-vous authentifier" && r.URL.Path != "/refresh" {
                    // Récupérer le refresh token
                    refreshCookie, err := r.Cookie("refresh_token")
                    if err != nil || refreshCookie.Value == "" {
                        http.Redirect(w, r, "/login", http.StatusSeeOther)
                        return
                    }

                    refreshToken := refreshCookie.Value

                    // Valider le refresh token
                    refreshClaims, err := Utils.ValidateJWT(refreshToken)
                    if err != nil || refreshClaims["type"] != "refresh" {
                        // Refresh token invalide ou expiré, supprimer les cookies et rediriger vers login
                        Utils.ClearTokenCookies(w)

                        http.Redirect(w, r, "/login", http.StatusSeeOther)
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
                        http.Error(w, "La génération de l'access token a échoué", http.StatusInternalServerError)
                        return
                    }

                    // Générer un nouveau refresh token pour la rotation
                    newRefreshToken, err := Utils.GenerateRefreshToken(userID)
                    if err != nil {
                        http.Error(w, "La génération du refresh token a échoué", http.StatusInternalServerError)
                        return
                    }

                    // Mettre à jour les cookies avec les nouveaux tokens
                    Utils.SetAccessTokenCookie(w, newAccessToken)
                    Utils.SetRefreshTokenCookie(w, newRefreshToken)

                    // Continuer avec la requête en injectant le nouveau access token
                } else {
                    // Autre type d'erreur de validation du token
                    http.Error(w, "Token invalide", http.StatusUnauthorized)
                    return
                }
            } else {
                // Vous pouvez ajouter les informations de l'utilisateur au contexte si nécessaire
                // par exemple:
                // ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])
                // next.ServeHTTP(w, r.WithContext(ctx))
            }

            next.ServeHTTP(w, r)
        })
    }
}