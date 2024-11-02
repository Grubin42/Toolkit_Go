package Utils

import (
    "html/template"
    "log"
    "path/filepath"
)

func LoadTemplates(specificViews ...string) *template.Template {
    templatePaths := []string{
        filepath.Join("cmd", "Presentation", "Views", "Layout", "base.html"),
        filepath.Join("cmd", "Presentation", "Views", "Layout", "navbar.html"),
    }

    for _, view := range specificViews {
        fullPath := filepath.Join("cmd", "Presentation", "Views", view)
        log.Printf("Chargement du template : %s\n", fullPath)
        templatePaths = append(templatePaths, fullPath)
    }

    tmpl, err := template.ParseFiles(templatePaths...)
    if err != nil {
        log.Fatalf("Erreur lors du chargement des templates : %v", err)
    }

    return tmpl
}