package Controllers

import (
	"html/template"
	"log"
    "fmt"
	"net/http"
	"path/filepath"
)

type RegisterController struct {
    templates *template.Template
}

func NewRegisterController() *RegisterController {
    // Charger les templates au démarrage
    tmpl, err := template.ParseFiles(
        filepath.Join("cmd", "Presentation", "Views", "Layout", "base.html"),
        filepath.Join("cmd", "Presentation", "Views", "Register", "index.html"),
    )
    if err != nil {
        log.Fatalf("Erreur lors du parsing des templates: %v", err)
    }

    return &RegisterController{
        templates: tmpl,
    }
}

func (hc *RegisterController) HandleIndex(w http.ResponseWriter, r *http.Request) {
    // Vérifiez la méthode HTTP
    if r.Method == http.MethodPost {
        // Récupérer les données du formulaire
        username := r.FormValue("username")
        email := r.FormValue("email")
        password := r.FormValue("password")
        confirmPassword := r.FormValue("confirm_password")

        // Ici, vous pouvez ajouter la logique pour traiter les données
        // Par exemple, valider les champs ou enregistrer l'utilisateur dans une base de données

        // Pour l'instant, nous affichons les valeurs récupérées dans les logs
        fmt.Fprintf(w,"Username: %s, Email: %s, Password: %s, Confirm Password: %s", username, email, password, confirmPassword)
		
        // Rediriger ou afficher un message de succès (à personnaliser selon les besoins)
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    // Si la méthode est GET, afficher le formulaire d'inscription
    data := struct {
        Title string
    }{
        Title: "register",
    }

    // Exécuter le template 'base.html' en injectant 'Register/index.html'
    err := hc.templates.ExecuteTemplate(w, "base", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
