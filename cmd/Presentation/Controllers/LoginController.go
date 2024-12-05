// cmd/Presentation/Controllers/LoginController.go
package Controllers

import (
    "database/sql"
    "net/http"
    "errors"
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
            lc.HandleError(w, r, "Connexion", "Erreur lors de la soumission du formulaire.", nil, http.StatusBadRequest)
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
            lc.HandleError(w, r, "Connexion", "Validation des champs échouée.", fieldErrors, http.StatusBadRequest)
            return
        }

        userID, err := lc.loginService.LoginUser(username, password)
        if err != nil {
            switch {
            case errors.Is(err, Errors.ErrUserNotFound):
                fieldErrors["username"] = "Ce nom d'utilisateur ou email est incorrect."
                lc.HandleError(w, r, "Connexion", "Validation des champs échouée.", fieldErrors, http.StatusUnauthorized)
            case errors.Is(err, Errors.ErrInvalidPassword):
                fieldErrors["password"] = "Le mot de passe est incorrect."
                lc.HandleError(w, r, "Connexion", "Validation des champs échouée.", fieldErrors, http.StatusUnauthorized)
            default:
                lc.HandleError(w, r, "Connexion", "Une erreur est survenue.", nil, http.StatusInternalServerError)
            }
            return
        }

        // Générer les tokens
        accessToken, refreshToken, err := Utils.GenerateTokens(userID)
        if err != nil {
            lc.HandleError(w, r, "Connexion", "Erreur lors de la génération du token.", nil, http.StatusInternalServerError)
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
    lc.Render(w, r, map[string]interface{}{
        "Title": "Connexion",
    })
}