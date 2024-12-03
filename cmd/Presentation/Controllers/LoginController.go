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
    tmpl := Utils.LoadTemplates(
        "Login/index.html",
        "Login/Component/loginForm.html",
    )

    return &LoginController{
        BaseController{
            Templates: tmpl,
        },
        Services.NewLoginService(db),
    }
}

func (lc *LoginController) HandleIndex(w http.ResponseWriter, r *http.Request) {
    errorsMap := make(map[string]string)

    if r.Method == http.MethodPost {
        if err := r.ParseForm(); err != nil {
            Utils.WriteError(w, http.StatusBadRequest, "Erreur lors de la soumission du formulaire.", nil)
            return
        }

        username := r.FormValue("username")
        password := r.FormValue("password")

        if username == "" {
            errorsMap["username"] = "Le nom d'utilisateur ou l'email est requis."
        }
        if password == "" {
            errorsMap["password"] = "Le mot de passe est requis."
        }

        if len(errorsMap) > 0 {
            Utils.WriteError(w, http.StatusBadRequest, "Validation des champs échouée.", errorsMap)
            return
        }

        _, userID, err := lc.loginService.LoginUser(username, password)
        if err != nil {
            if err.Error() == "username_not_found" {
                errorsMap["username"] = "Ce nom d'utilisateur ou email est incorrect."
            } else if err.Error() == "incorrect_password" {
                errorsMap["password"] = "Le mot de passe est incorrect."
            } else {
                Utils.WriteError(w, http.StatusInternalServerError, "Échec de la connexion. Veuillez vérifier vos identifiants.", nil)
                return
            }
            Utils.WriteError(w, http.StatusBadRequest, "Validation des champs échouée.", errorsMap)
            return
        }

        // Générer les tokens
        accessToken, refreshToken, err := Utils.GenerateTokens(userID)
        if err != nil {
            Utils.WriteError(w, http.StatusInternalServerError, Errors.ErrorTokenGeneration, nil)
            return
        }

        // Stocker les tokens dans des cookies sécurisés
        Utils.SetAccessTokenCookie(w, accessToken)
        Utils.SetRefreshTokenCookie(w, refreshToken)

        // Rediriger vers la page d'accueil après une connexion réussie
        w.Header().Set("HX-Redirect", "/")
        return
    }

    // Préparer les données spécifiques pour une requête GET sans erreurs
    lc.HandleError(w, r, "Connexion", "", nil)
}