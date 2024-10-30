// cmd/Presentation/Controllers/base_controller.go
package Controllers

import (
    "html/template"
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
)

// BaseController contient les templates et fournit une méthode pour rendre les templates avec des données communes
type BaseController struct {
    Templates *template.Template
}

// Render exécute le template avec les données spécifiques et injecte IsAuthenticated
func (bc *BaseController) Render(w http.ResponseWriter, r *http.Request, specificData interface{}) {
    // Créer un map pour contenir toutes les données passées au template
    data := map[string]interface{}{
        "IsAuthenticated": Utils.IsAuthenticated(r),
    }

    // Si des données spécifiques sont fournies, les ajouter au map
    if specificData != nil {
        if sd, ok := specificData.(map[string]interface{}); ok {
            for key, value := range sd {
                data[key] = value
            }
        } else {
            data["Data"] = specificData
        }
    }

    // Exécuter le template de base avec les données
    err := bc.Templates.ExecuteTemplate(w, "base", data)
    if err != nil {
        http.Error(w, "Erreur de template", http.StatusInternalServerError)
    }
}