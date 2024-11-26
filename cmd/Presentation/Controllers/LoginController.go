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
        "Login/Component/loginForm.html",
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
    errors := make(map[string]string)

    if r.Method == http.MethodPost {
        // Analyse des données du formulaire
        if err := r.ParseForm(); err != nil {
            lc.HandleError(w, r, "Connexion", "Erreur lors de la soumission du formulaire.", nil)
            return
        }

        // Récupération des valeurs des champs
        identifier := r.FormValue("username")
        password := r.FormValue("password")

        // Validation des champs d'input
        if identifier == "" {
            errors["username"] = "Le nom d'utilisateur ou l'email est requis."
        }
        if password == "" {
            errors["password"] = "Le mot de passe est requis."
        }

        // Si des erreurs d'input existent, les renvoyer avant d'appeler LoginUser
        if len(errors) > 0 {
            lc.HandleError(w, r, "Connexion", "", errors)
            return
        }

        // Tentative de connexion de l'utilisateur
        _, userID, err := lc.loginService.LoginUser(identifier, password)
        if err != nil {
            // Gestion des erreurs spécifiques à chaque champ
            if err.Error() == "username_not_found" {
                errors["username"] = "Ce nom d'utilisateur ou email est incorrect."
            } else if err.Error() == "incorrect_password" {
                errors["password"] = "Le mot de passe est incorrect."
            } else {
                // Si l'erreur n'est pas spécifiée, on affiche un message général
                lc.HandleError(w, r, "Connexion", "Échec de la connexion. Veuillez vérifier vos identifiants.", nil)
                return
            }
            // Affiche les erreurs spécifiques au champ
            lc.HandleError(w, r, "Connexion", "", errors)
            return
        }

        // Révoquer tous les anciens refresh tokens
        err = lc.refreshService.RevokeAllRefreshTokens(userID)
        if err != nil {
            lc.HandleError(w, r, "Connexion", Errors.ErrorRevokeAllTokens, nil)
            return
        }

        // Générer les tokens
        accessToken, err := Utils.GenerateAccessToken(userID)
        if err != nil {
            lc.HandleError(w, r, "Connexion", Errors.ErrorTokenGeneration, nil)
            return
        }

        refreshToken, err := Utils.GenerateRefreshToken(userID)
        if err != nil {
            lc.HandleError(w, r, "Connexion", Errors.ErrorRefreshTokenGeneration, nil)
            return
        }

        // Enregistrer le refresh token dans la base de données
        expiresAt := time.Now().Add(Utils.GetRefreshTokenExpiration())
        err = lc.refreshService.SaveRefreshToken(userID, refreshToken, expiresAt)
        if err != nil {
            lc.HandleError(w, r, "Connexion", Errors.ErrorRefreshTokenSave, nil)
            return
        }

        // Stocker les tokens dans des cookies sécurisés
        http.SetCookie(w, &http.Cookie{
            Name:     "jwt_token",
            Value:    accessToken,
            Expires:  time.Now().Add(Utils.GetAccessTokenExpiration()),
            HttpOnly: true,
            Secure:   false, // À activer en production (HTTPS)
            SameSite: http.SameSiteStrictMode,
            Path:     "/",
        })

        http.SetCookie(w, &http.Cookie{
            Name:     "refresh_token",
            Value:    refreshToken,
            Expires:  expiresAt,
            HttpOnly: true,
            Secure:   false, // À activer en production (HTTPS)
            SameSite: http.SameSiteStrictMode,
            Path:     "/refresh",
        })

        // Rediriger vers la page d'accueil après une connexion réussie
        w.Header().Set("HX-Redirect", "/")
        return
    }

    // Préparer les données spécifiques pour une requête GET sans erreurs
    lc.HandleError(w, r, "Connexion", "", nil)
}