package Controllers

import (
    "net/http"
    "database/sql"
    "time"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Services"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
)

type LoginController struct {
    BaseController
    loginService *Services.LoginService
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
    }
}

func (lc *LoginController) HandleIndex(w http.ResponseWriter, r *http.Request) {
    var errorMessage string

    if r.Method == http.MethodPost {
        identifier := r.FormValue("username")
        password := r.FormValue("password")

        // Appeler le service de connexion pour obtenir le JWT
        status, token, err := lc.loginService.LoginUser(identifier, password)
        if err != nil {
            w.WriteHeader(status)
            errorMessage = err.Error()
            http.Error(w, errorMessage, status)
            return
        }
        // Stocker le JWT dans un cookie sécurisé
        http.SetCookie(w, &http.Cookie{
            Name:     "jwt_token",
            Value:    token,
            Expires:  time.Now().Add(Utils.GetTokenExpiration()),
            HttpOnly: true, // Pour empêcher l'accès côté client
            Secure:   false, // À activer en production (HTTPS)
            SameSite: http.SameSiteStrictMode,  // Empêche les attaques CSRF
        })
        w.Header().Set("HX-Redirect", "/")
    }

    // Préparer les données spécifiques à la vue
    specificData := map[string]interface{}{
        "Title":           "Connexion",
        "ErrorMessage":    errorMessage,
    }

    // Utiliser la méthode Render du BaseController
    lc.Render(w, r, specificData)
}