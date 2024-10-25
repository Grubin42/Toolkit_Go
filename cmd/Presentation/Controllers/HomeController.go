package Controllers

import (
    "html/template"
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
)

type HomeController struct {
    templates *template.Template
}

func NewHomeController() *HomeController {
    return &HomeController{
        templates: Utils.LoadTemplates("Home/index.html"),
    }
}

func (hc *HomeController) HandleIndex(w http.ResponseWriter, r *http.Request) {
    // Vérifier la présence du cookie JWT pour déterminer si l'utilisateur est connecté
    isAuthenticated := Utils.IsAuthentificated(r)

    // Passer les données au template
    data := struct {
        Title           string
        IsAuthenticated bool
    }{
        Title:           "Accueil",
        IsAuthenticated: isAuthenticated,
    }

    // Exécuter le template 'base.html' en injectant 'Home/index.html'
    err := hc.templates.ExecuteTemplate(w, "base", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}