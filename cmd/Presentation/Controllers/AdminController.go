// cmd/Presentation/Controllers/AdminController.go
package Controllers

import (
    "html/template"
    "log"
    "net/http"
    "path/filepath"
)

type AdminController struct {
    templates *template.Template
}

func NewAdminController() *AdminController {
    // Charger les templates nécessaires
    tmpl, err := template.ParseFiles(
        filepath.Join("cmd", "Presentation", "Views", "Layout", "base.html"),
        filepath.Join("cmd", "Presentation", "Views", "Layout", "navbar.html"),
        filepath.Join("cmd", "Presentation", "Views", "Admin", "index.html"),
    )
    if err != nil {
        log.Fatalf("Erreur lors du parsing des templates Admin: %v", err)
    }

    return &AdminController{
        templates: tmpl,
    }
}

func (ac *AdminController) HandleIndex(w http.ResponseWriter, r *http.Request) {
    data := struct {
        Title string
        // Ajoutez d'autres champs si nécessaire
    }{
        Title: "Administration",
    }

    // Exécuter le template 'base' avec les données
    err := ac.templates.ExecuteTemplate(w, "base", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}