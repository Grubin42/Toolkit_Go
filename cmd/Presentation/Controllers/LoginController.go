package Controllers

import (
    "html/template"
    "net/http"
    "database/sql"
    "time"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Services"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
)

type LoginController struct {
    templates    *template.Template
    loginService *Services.LoginService
}

func NewLoginController(db *sql.DB) *LoginController {
    return &LoginController{
        templates:    Utils.LoadTemplates("Login/index.html"),
        loginService: Services.NewLoginService(db),
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
            Expires:  time.Now().Add(24 * time.Hour),
            HttpOnly: true,
            Secure:   false,
            Path:     "/",
            SameSite: http.SameSiteStrictMode,
        })

        // Rediriger après une connexion réussie
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    // Passer l'état de connexion à la vue
    data := struct {
        Title           string
        ErrorMessage    string
        IsAuthenticated bool
    }{
        Title:           "Accueil",
        ErrorMessage:    errorMessage,
        IsAuthenticated: isAuthenticated,
    }

    err := lc.templates.ExecuteTemplate(w, "base", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}