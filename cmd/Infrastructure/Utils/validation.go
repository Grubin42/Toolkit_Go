package Utils

import (
    "errors"
    "regexp"
    "unicode"
	"strings"
)

// ValidateUsername vérifie que le username a entre 3 et 20 caractères
func ValidateUsername(username string) error {
    if len(username) < 3 || len(username) > 20 {
        return errors.New("le nom d'utilisateur doit contenir entre 3 et 20 caractères")
    }
    return nil
}

// ValidateEmail vérifie que l'email a un format valide
func ValidateEmail(email string) error {
    emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
    if !emailRegex.MatchString(email) {
        return errors.New("l'adresse email n'est pas valide")
    }
    return nil
}

// ValidatePassword vérifie que le mot de passe contient au moins 8 caractères,
// une majuscule, une minuscule, un chiffre, et un caractère spécial
func ValidatePassword(password string) []string {
    var errors []string

    // Vérifier la longueur du mot de passe
    if len(password) < 8 {
        errors = append(errors, "le mot de passe doit contenir au moins 8 caractères")
    }
    if !containsUppercase(password) {
        errors = append(errors, "le mot de passe doit contenir au moins une majuscule")
    }
    if !containsLowercase(password) {
        errors = append(errors, "le mot de passe doit contenir au moins une minuscule")
    }
    if !containsDigit(password) {
        errors = append(errors, "le mot de passe doit contenir au moins un chiffre")
    }
    if !containsSpecialChar(password) {
        errors = append(errors, "le mot de passe doit contenir au moins un caractère spécial")
    }

    return errors
}

// containsUppercase vérifie si une chaîne contient une majuscule
func containsUppercase(s string) bool {
    for _, r := range s {
        if unicode.IsUpper(r) {
            return true
        }
    }
    return false
}

// containsLowercase vérifie si une chaîne contient une minuscule
func containsLowercase(s string) bool {
    for _, r := range s {
        if unicode.IsLower(r) {
            return true
        }
    }
    return false
}

// containsDigit vérifie si une chaîne contient un chiffre
func containsDigit(s string) bool {
    for _, r := range s {
        if unicode.IsDigit(r) {
            return true
        }
    }
    return false
}

// containsSpecialChar vérifie si une chaîne contient un caractère spécial
func containsSpecialChar(s string) bool {
    specialChars := "!@#$%^&*()_-+={}[]|:;'<>,.?/~`"
    for _, r := range s {
        if strings.ContainsRune(specialChars, r) {
            return true
        }
    }
    return false
}