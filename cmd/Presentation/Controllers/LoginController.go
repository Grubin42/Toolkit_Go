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

    // Vérifier la présence du cookie JWT pour déterminer si l'utilisateur est connecté
    isAuthenticated := Utils.IsAuthentificated(r)
    
    if r.Method == http.MethodPost {
        identifier := r.FormValue("username")
        password := r.FormValue("password")

        // Appeler le service de connexion pour obtenir le JWT
        status, token, err := lc.loginService.LoginUser(identifier, password)
        if err != nil {
            http.Error(w, err.Error(), status)
            return
        }

        // Stocker le JWT dans un cookie sécurisé
        http.SetCookie(w, &http.Cookie{
            Name:     "jwt_token",
            Value:    token,
<<<<<<< HEAD
            Expires:  time.Now().Add(24 * time.Hour),
            HttpOnly: true,
            Secure:   false,
            Path:     "/",
            SameSite: http.SameSiteStrictMode,
=======
            Expires:  time.Now().Add(Utils.GetTokenExpiration()),
            HttpOnly: true, // Pour empêcher l'accès côté client
            Secure:   false, // À activer en production (HTTPS)
            SameSite: http.SameSiteStrictMode,  // Empêche les attaques CSRF
>>>>>>> origin/gael-dev
        })

        // Rediriger après une connexion réussie
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

<<<<<<< HEAD
    // Passer l'état de connexion à la vue
    data := struct {
        Title           string
        ErrorMessage    string
        IsAuthenticated bool
    }{
        Title:           "Accueil",
        ErrorMessage:    errorMessage,
        IsAuthenticated: isAuthenticated,
=======
    // Préparer les données spécifiques à la vue
    specificData := map[string]interface{}{
        "Title":           "Connexion",
        "ErrorMessage":    errorMessage,
>>>>>>> origin/gael-dev
    }

    // Utiliser la méthode Render du BaseController
    lc.Render(w, r, specificData)
}