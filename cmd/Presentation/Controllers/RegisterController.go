package Controllers

import (
	"html/template"
	"log"
    "fmt"
	"net/http"
	"path/filepath"
    "database/sql"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Services"
)

type RegisterController struct {
    templates     *template.Template
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

func (hc *RegisterController) HandleIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        email := r.FormValue("email")
        password := r.FormValue("password")
        confirmPassword := r.FormValue("confirm_password")

        // Appeler le service d'enregistrement
        err := hc.registerService.RegisterUser(username, email, password, confirmPassword)
        if err != nil {
            // Gérer l'erreur (ex: afficher un message d'erreur à l'utilisateur)
            fmt.Fprintf(w, "Erreur lors de l'enregistrement : %s", err.Error())
            return
        }

        // Redirection après l'enregistrement réussi
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    data := struct {
        Title string
    }{
        Title: "register",
    }

    err := hc.templates.ExecuteTemplate(w, "base", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}