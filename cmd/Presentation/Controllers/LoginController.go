package Controllers

import (
    "html/template"
    "log"
    "net/http"
    "path/filepath"
    "database/sql"
    "time"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Services"
)

type LoginController struct {
    templates    *template.Template
    loginService *Services.LoginService
}

func NewLoginController(db *sql.DB) *LoginController {
    tmpl, err := template.ParseFiles(
        filepath.Join("cmd", "Presentation", "Views", "Layout", "base.html"),
        filepath.Join("cmd", "Presentation", "Views", "Layout", "navbar.html"),
        filepath.Join("cmd", "Presentation", "Views", "Login", "index.html"),
    )
    if err != nil {
        log.Fatalf("Erreur lors du parsing des templates: %v", err)
    }

    return &LoginController{
        templates:    tmpl,
        loginService: Services.NewLoginService(db),
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
            Expires:  time.Now().Add(24 * time.Hour),
            HttpOnly: true, // Pour empêcher l'accès côté client
            Secure:   false, // À activer en production (HTTPS)
            SameSite: http.SameSiteStrictMode,  // Empêche les attaques CSRF
        })

        // Rediriger après une connexion réussie
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    data := struct {
        Title       string
        ErrorMessage string
    }{
        Title:       "login",
        ErrorMessage: errorMessage,
    }

    err := lc.templates.ExecuteTemplate(w, "base", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}