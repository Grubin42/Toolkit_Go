package Controllers

import (
    "net/http"
    "database/sql"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Services"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
)

type RegisterController struct {
    BaseController
    registerService *Services.RegisterService
}

func NewRegisterController(db *sql.DB) *RegisterController {
    // Charger les templates partagés et spécifique
    return &RegisterController{
        BaseController{
            Templates: Utils.LoadTemplates(
                "Register/index.html",
            ),
        },
        Services.NewRegisterService(db),
    }
}

// HandleIndex gère la route d'enregistrement et affiche les erreurs ou les succès
func (rc *RegisterController) HandleIndex(w http.ResponseWriter, r *http.Request) {
    var errorMessage string

    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        email := r.FormValue("email")
        password := r.FormValue("password")
        confirmPassword := r.FormValue("confirm_password")

        // Appeler le service d'enregistrement et récupérer le code HTTP
        status, err := rc.registerService.RegisterUser(username, email, password, confirmPassword)
        if err != nil {
            w.WriteHeader(status)  // Écrire le bon code HTTP
            errorMessage = err.Error()  // Capturer l'erreur pour l'envoyer à la vue
            rc.HandleError(w, r, "Inscription", errorMessage)
            return
        }
        if Utils.IsHtmxRequest(r) {
            // Vous pouvez choisir de rediriger ou de montrer un message de succès
            // Ici, nous utilisons HX-Redirect pour rediriger vers la page de login
            w.Header().Set("HX-Redirect", "/login")
            return
        }
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    // Préparer les données spécifiques à la vue
    specificData := map[string]interface{}{
        "Title":           "Inscription",
        "ErrorMessage":    errorMessage,
    }

    // Utiliser la méthode Render du BaseController
    rc.Render(w, r, specificData)
}