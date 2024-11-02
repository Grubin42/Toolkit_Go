// cmd/Infrastructure/Middleware/AuthMiddleware.go
package Middleware

import (
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Services"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Errors"
    "time"
    "database/sql"
)

func AuthMiddleware(db *sql.DB) func(http.Handler) http.Handler {
    refreshService := Services.NewRefreshService(db)

    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Récupérer le access token depuis le cookie
            cookie, err := r.Cookie("jwt_token")
            if err != nil {
                if err == http.ErrNoCookie {
                    http.Redirect(w, r, "/login", http.StatusSeeOther)
                    return
                }
                http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
                return
            }

            accessToken := cookie.Value
            if accessToken == "" {
                http.Redirect(w, r, "/login", http.StatusSeeOther)
                return
            }

            // Valider le access token
            _ , err = Utils.ValidateJWT(accessToken)
            if err != nil {
                // Si le token est expiré, tenter de rafraîchir
                // Vérifier si l'erreur est due à l'expiration
                if err.Error() == "token invalide ou expiré" {
                    // Récupérer le refresh token
                    refreshCookie, err := r.Cookie("refresh_token")
                    if err != nil || refreshCookie.Value == "" {
                        http.Redirect(w, r, "/login", http.StatusSeeOther)
                        return
                    }

                    refreshToken := refreshCookie.Value

                    // Valider le refresh token
                    userID, err := refreshService.ValidateRefreshToken(refreshToken)
                    if err != nil {
                        // Refresh token invalide ou expiré, supprimer les cookies et rediriger vers login
                        http.SetCookie(w, &http.Cookie{
                            Name:     "jwt_token",
                            Value:    "",
                            Expires:  time.Unix(0, 0),
                            HttpOnly: true,
                            Secure:   true,
                            SameSite: http.SameSiteStrictMode,
                            Path:     "/",
                        })

                        http.SetCookie(w, &http.Cookie{
                            Name:     "refresh_token",
                            Value:    "",
                            Expires:  time.Unix(0, 0),
                            HttpOnly: true,
                            Secure:   true,
                            SameSite: http.SameSiteStrictMode,
                            Path:     "/refresh",
                        })

                        http.Redirect(w, r, "/login", http.StatusSeeOther)
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
                    err = refreshService.SaveRefreshToken(userID, newRefreshToken, expiresAt)
                    if err != nil {
                        http.Error(w, Errors.ErrorRefreshTokenSave, http.StatusInternalServerError)
                        return
                    }

                    // Révoquer l'ancien refresh token
                    err = refreshService.RevokeRefreshToken(refreshToken)
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

                    // Continuer avec la requête en injectant le nouveau access token
                } else {
                    // Autre type d'erreur de validation du token
                    http.Error(w, "Token invalide", http.StatusUnauthorized)
                    return
                }
            }

            // Vous pouvez ajouter les informations de l'utilisateur au contexte si nécessaire
            // par exemple:
            // ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])
            // next.ServeHTTP(w, r.WithContext(ctx))

            next.ServeHTTP(w, r)
        })
    }
}