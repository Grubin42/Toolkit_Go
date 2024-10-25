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
        templates: Utils.LoadTemplates("Home/index.html"),//exemple si plusieur templates/composants "templates: Utils.LoadTemplates("Home/index.html", "Home/featured.html", "Home/contact.html"),"
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