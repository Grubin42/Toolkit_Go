// cmd/Presentation/Controllers/AdminController.go
package Controllers

import (
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
)

type AdminController struct {
    BaseController
}

func NewAdminController() *AdminController {
    return &AdminController{
        BaseController{
            Templates: Utils.LoadTemplates( "Admin/index.html"),
        },
    }
}

func (ac *AdminController) HandleIndex(w http.ResponseWriter, r *http.Request) {

    // Préparer les données spécifiques à la vue
    specificData := map[string]interface{}{
        "Title": "Administration",
        // Ajoutez d'autres données spécifiques si nécessaire
    }

    // Utiliser la méthode Render du BaseController
    ac.Render(w, r, specificData)
}