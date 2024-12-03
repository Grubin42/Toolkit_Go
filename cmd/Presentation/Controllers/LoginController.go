// cmd/Presentation/Controllers/LoginController.go
package Controllers

import (
    "database/sql"
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Services"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Errors"
)

type LoginController struct {
    BaseController
    loginService *Services.LoginService
}

func NewLoginController(db *sql.DB) *LoginController {

    return &LoginController{
        BaseController{
            Templates: Utils.LoadTemplates(
                "Login/index.html",
                "Login/Component/loginForm.html",
            ),
        },
        Services.NewLoginService(db),
    }
}

func (lc *LoginController) HandleIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        if err := r.ParseForm(); err != nil {
            lc.HandleError(w, r, "Connexion", "Erreur lors de la soumission du formulaire.", nil, Errors.ServerError)
            return
        }

        username := r.FormValue("username")
        password := r.FormValue("password")
        fieldErrors := make(map[string]string)

        if username == "" {
            fieldErrors["username"] = "Le nom d'utilisateur ou l'email est requis."
        }
        if password == "" {
            fieldErrors["password"] = "Le mot de passe est requis."
        }

        if len(fieldErrors) > 0 {
            lc.HandleError(w, r, "Connexion", "Validation des champs échouée.", fieldErrors, Errors.ValidationError)
            return
        }

        _, userID, err := lc.loginService.LoginUser(username, password)
        if err != nil {
            fieldErrors := make(map[string]string)
            switch e := err.(type) {
            case *Errors.AppError:
                if e.Type == Errors.AuthenticationError {
                    if e.Message == "username_not_found" {
                        fieldErrors["username"] = "Ce nom d'utilisateur ou email est incorrect."
                    } else if e.Message == "incorrect_password" {
                        fieldErrors["password"] = "Le mot de passe est incorrect."
                    }
                    lc.HandleError(w, r, "Connexion", "Validation des champs échouée.", fieldErrors, Errors.ValidationError)
                    return
                } else {
                    lc.HandleError(w, r, "Connexion", e.Message, nil, Errors.ServerError)
                    return
                }
            default:
                lc.HandleError(w, r, "Connexion", Errors.ErrorLoginFailed, nil, Errors.ServerError)
                return
            }
        }

        // Générer les tokens
        accessToken, refreshToken, err := Utils.GenerateTokens(userID)
        if err != nil {
            lc.HandleError(w, r, "Connexion", Errors.ErrorTokenGeneration, nil, Errors.ServerError)
            return
        }

        // Stocker les tokens dans des cookies sécurisés
        Utils.SetAccessTokenCookie(w, accessToken)
        Utils.SetRefreshTokenCookie(w, refreshToken)

        // Rediriger via HTMX
        w.Header().Set("HX-Redirect", "/")
        return
    }

    // Préparer les données spécifiques pour une requête GET sans erreurs
    lc.HandleError(w, r, "Connexion", "", nil, Errors.ServerError)
}