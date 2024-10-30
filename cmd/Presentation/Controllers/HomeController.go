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
<<<<<<< HEAD
    // Vérifier la présence du cookie JWT pour déterminer si l'utilisateur est connecté
    isAuthenticated := Utils.IsAuthentificated(r)

    // Passer les données au template
    data := struct {
        Title           string
        IsAuthenticated bool
    }{
        Title:           "Accueil",
        IsAuthenticated: isAuthenticated,
=======

    // Préparer les données spécifiques à la vue
    specificData := map[string]interface{}{
        "Title": "Accueil",
        // Ajoutez d'autres données spécifiques si nécessaire
>>>>>>> origin/gael-dev
    }

    // Utiliser la méthode Render du BaseController
    hc.Render(w, r, specificData)
}