package Controllers

import (
    "html/template"
    "log"
    "net/http"
    "path/filepath"
    "database/sql"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Services"
)

type RegisterController struct {
    templates       *template.Template
    registerService *Services.RegisterService
}

func NewRegisterController(db *sql.DB) *RegisterController {
    tmpl, err := template.ParseFiles(
        filepath.Join("cmd", "Presentation", "Views", "Layout", "base.html"),
        filepath.Join("cmd", "Presentation", "Views", "Register", "index.html"),
    )
    if err != nil {
        log.Fatalf("Erreur lors du parsing des templates: %v", err)
    }

    return &RegisterController{
        templates:       tmpl,
        registerService: Services.NewRegisterService(db),
    }
}

// HandleIndex gère la route d'enregistrement et affiche les erreurs ou les succès
func (rc *RegisterController) HandleIndex(w http.ResponseWriter, r *http.Request) {
    var errorMessage string

    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        email := r.FormValue("email")
        password := r.FormValue("password")
        confirmPassword := r.FormValue("confirm_password")

        // Appeler le service d'enregistrement et récupérer le code HTTP
        status, err := rc.registerService.RegisterUser(username, email, password, confirmPassword)
        if err != nil {
            w.WriteHeader(status)  // Écrire le bon code HTTP
            errorMessage = err.Error()  // Capturer l'erreur pour l'envoyer à la vue
        } else {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }
    }

    data := struct {
        Title       string
        ErrorMessage string
    }{
        Title:       "Register",
        ErrorMessage: errorMessage,  // Passer le message d'erreur à la vue
    }

    err := rc.templates.ExecuteTemplate(w, "base", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}