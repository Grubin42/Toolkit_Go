package Controllers

import (
    "net/http"
    "database/sql"
    "time"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Services"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Errors"
)

type LoginController struct {
    BaseController
    loginService *Services.LoginService
    refreshService *Services.RefreshService
}

func NewLoginController(db *sql.DB) *LoginController {
    // Charger les templates partagés et spécifiques
    tmpl := Utils.LoadTemplates(
        "Login/index.html",
    )

    return &LoginController{
        BaseController{
            Templates: tmpl,
        },
        Services.NewLoginService(db),
        Services.NewRefreshService(db),
    }
}

func (lc *LoginController) HandleIndex(w http.ResponseWriter, r *http.Request) {
    var errorMessage string

    if r.Method == http.MethodPost {
        identifier := r.FormValue("username")
        password := r.FormValue("password")

        // Appeler le service de connexion pour obtenir l'ID de l'utilisateur
        status, userID, err := lc.loginService.LoginUser(identifier, password)
        if err != nil {
            w.WriteHeader(status)
            errorMessage = err.Error()
            lc.HandleError(w, r, "Connexion", errorMessage)
            return
        }


        // Révoquer tous les anciens refresh tokens
        err = lc.refreshService.RevokeAllRefreshTokens(userID)
        if err != nil {
            errorMessage = Errors.ErrorRevokeAllTokens
            lc.HandleError(w, r, "Connexion", errorMessage)
            return
        }


        // Générer les tokens
        accessToken, err := Utils.GenerateAccessToken(userID)
        if err != nil {
            errorMessage = Errors.ErrorTokenGeneration
            lc.HandleError(w, r, "Connexion", errorMessage)
            return
        }

        refreshToken, err := Utils.GenerateRefreshToken(userID)
        if err != nil {
            errorMessage = Errors.ErrorRefreshTokenGeneration
            lc.HandleError(w, r, "Connexion", errorMessage)
            return
        }

        // Enregistrer le refresh token dans la base de données
        expiresAt := time.Now().Add(Utils.GetRefreshTokenExpiration())
        err = lc.refreshService.SaveRefreshToken(userID, refreshToken, expiresAt)
        if err != nil {
            errorMessage = Errors.ErrorRefreshTokenSave
            lc.HandleError(w, r, "Connexion", errorMessage)
            return
        }
        
        // Stocker les tokens dans des cookies sécurisés
        http.SetCookie(w, &http.Cookie{
            Name:     "jwt_token",
            Value:    accessToken,
            Expires:  time.Now().Add(Utils.GetAccessTokenExpiration()),
            HttpOnly: true, // Pour empêcher l'accès côté client
            Secure:   false, // À activer en production (HTTPS)
            SameSite: http.SameSiteStrictMode,  // Empêche les attaques CSRF
            Path:     "/",
        })

        http.SetCookie(w, &http.Cookie{
            Name:     "refresh_token",
            Value:    refreshToken,
            Expires:  expiresAt,
            HttpOnly: true,
            Secure:   false, // À activer en production (HTTPS)
            SameSite: http.SameSiteStrictMode,
            Path:     "/refresh", // Définir un chemin spécifique pour le refresh
        })
        w.Header().Set("HX-Redirect", "/")
        return
    }

    // Préparer les données spécifiques à la vue
    specificData := map[string]interface{}{
        "Title":           "Connexion",
        "ErrorMessage":    errorMessage,
    }

    // Utiliser la méthode Render du BaseController
    lc.Render(w, r, specificData)
}