package Utils

import (
    "html/template"
    "log"
    "path/filepath"
)

// LoadTemplates charge les templates communs ainsi qu'un ou plusieurs templates spécifiques
func LoadTemplates(specificViews ...string) *template.Template {
    // Définir une slice pour stocker tous les chemins des fichiers à charger
    templatePaths := []string{
        filepath.Join("cmd", "Presentation", "Views", "Layout", "base.html"),
        filepath.Join("cmd", "Presentation", "Views", "Layout", "navbar.html"),
    }

    // Ajouter les templates spécifiques passés en paramètre
    for _, view := range specificViews {
        templatePaths = append(templatePaths, filepath.Join("cmd", "Presentation", "Views", view))
    }

    // Charger tous les templates à partir de la liste des chemins
    tmpl, err := template.ParseFiles(templatePaths...)
    if err != nil {
        log.Fatalf("Erreur lors du chargement des templates : %v", err)
    }

    return tmpl
}