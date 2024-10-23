// cmd/Presentation/Controllers/AdminController.go
package Controllers

import (
    "html/template"
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
)

type AdminController struct {
    templates *template.Template
}

func NewAdminController() *AdminController {
    return &AdminController{
        templates: Utils.LoadTemplates("Admin/index.html"),
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