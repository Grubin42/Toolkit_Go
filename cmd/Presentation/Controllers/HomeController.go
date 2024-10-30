package Controllers

import (
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
)

type HomeController struct {
    BaseController
}

// NewHomeController initialise un HomeController avec les templates nécessaires
func NewHomeController() *HomeController {
    return &HomeController{
        BaseController{
            Templates: Utils.LoadTemplates( "Home/index.html"),
        },
    }
}

func (hc *HomeController) HandleIndex(w http.ResponseWriter, r *http.Request) {

    // Préparer les données spécifiques à la vue
    specificData := map[string]interface{}{
        "Title": "Accueil",
        // Ajoutez d'autres données spécifiques si nécessaire
    }

    // Utiliser la méthode Render du BaseController
    hc.Render(w, r, specificData)
}