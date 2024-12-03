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
    if Utils.IsHtmxRequest(r) {
        // Exécuter uniquement le template spécifique (par exemple, "Home/index.html")
        err := bc.Templates.ExecuteTemplate(w, "content", data)
        if err != nil {
            http.Error(w, "Erreur de template", http.StatusInternalServerError)
        }
        return
    }
    
    // Exécuter le template de base avec les données
    err := bc.Templates.ExecuteTemplate(w, "base", data)
    if err != nil {
        http.Error(w, "Erreur de template", http.StatusInternalServerError)
    }
}


// HandleError gère la redirection avec un message d'erreur
func (bc *BaseController) HandleError(w http.ResponseWriter, r *http.Request, title string, errorMessage string, errors map[string]string) {
    specificData := map[string]interface{}{
        "Title":        title,
        "ErrorMessage": errorMessage, // Message d'erreur général
        "Errors":       errors,       // Map des erreurs spécifiques aux champs
    }
    bc.Render(w, r, specificData)
}