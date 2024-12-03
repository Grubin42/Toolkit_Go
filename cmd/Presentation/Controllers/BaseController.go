// cmd/Presentation/Controllers/BaseController.go
package Controllers

import (
    "html/template"
    "log"
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Errors"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Models"
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

        // Exécuter le template spécifique (par exemple, "content")
        err := bc.Templates.ExecuteTemplate(w, "content", data)
        if err != nil {
            log.Printf("Erreur d'exécution du template 'content' : %v", err)
            bc.RenderHTMXError(w, "Une erreur est survenue lors du traitement de votre demande.", Errors.ServerError)
            return
        }
        return
    }

    // Exécuter le template de base avec les données
    err := bc.Templates.ExecuteTemplate(w, "base", data)
    if err != nil {
        log.Printf("Erreur d'exécution du template 'base' : %v", err)
        bc.RenderHTMXError(w, "Une erreur est survenue lors du chargement de la page.", Errors.ServerError)
    }
}

// RenderHTMXError rend un fragment HTML contenant un message d'erreur pour HTMX
func (bc *BaseController) RenderHTMXError(w http.ResponseWriter, message string, errType Errors.ErrorType) {
    errorData := map[string]interface{}{
        "ErrorMessage": message,
        "ErrorType":    errType,
    }
    log.Printf("Rendu du template 'errorMessage' avec les données : %+v", errorData)
    err := bc.Templates.ExecuteTemplate(w, "errorMessage", errorData)
    if err != nil {
        log.Printf("Erreur d'exécution du template 'errorMessage' : %v", err)
        http.Error(w, "Erreur interne du serveur1", http.StatusInternalServerError)
    }
}

// HandleError gère la redirection avec un message d'erreur via HTMX ou global
func (bc *BaseController) HandleError(w http.ResponseWriter, r *http.Request, title string, errorMessage string, fieldErrors map[string]string, errType Errors.ErrorType) {
    formErrors := Models.FormErrors{
        ErrorMessage: errorMessage,
        FieldErrors:  fieldErrors,
    }

    specificData := map[string]interface{}{
        "Title":      title,
        "FormErrors": formErrors,
        "Data": map[string]string{
            "username": r.FormValue("username"), // Passer les valeurs précédentes si nécessaire
        },
        "ErrorType": errType,
    }

    bc.Render(w, r, specificData)
}