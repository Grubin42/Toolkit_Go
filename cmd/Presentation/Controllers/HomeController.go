package Controllers

import (
    "html/template"
    "net/http"
    "path/filepath"
	"log"
)

type HomeController struct {
    templates *template.Template
}

func NewHomeController() *HomeController {
    // Charger les templates au démarrage
    tmpl, err := template.ParseFiles(
        filepath.Join("cmd", "Presentation", "Views", "Layout", "base.html"),
        filepath.Join("cmd", "Presentation", "Views", "Home", "index.html"),
    )
    if err != nil {
        log.Fatalf("Erreur lors du parsing des templates: %v", err)
    }

    return &HomeController{
        templates: tmpl,
    }
}

func (hc *HomeController) HandleIndex(w http.ResponseWriter, r *http.Request) {
    data := struct {
        Title string
        // Ajoutez d'autres données si nécessaire
    }{
        Title: "Accueil",
    }

    // Exécuter le template 'base.html' en injectant 'Home/index.html'
    err := hc.templates.ExecuteTemplate(w, "base", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}