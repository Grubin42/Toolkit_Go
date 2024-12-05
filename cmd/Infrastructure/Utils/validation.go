package Utils

import (
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Errors"
    "regexp"
    "unicode"
	"strings"
)

// ValidateUsername vérifie que le username a entre 3 et 20 caractères
func ValidateUsername(username string) error {
    if len(username) < 3 || len(username) > 20 {
        return Errors.ErrUsernameInvalidLength
    }
    return nil
}

// ValidateEmail vérifie que l'email a un format valide
func ValidateEmail(email string) error {
    emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
    if !emailRegex.MatchString(email) {
        return Errors.ErrEmailInvalidFormat
    }
    return nil
}

// ValidatePassword vérifie que le mot de passe contient au moins 8 caractères,
// une majuscule, une minuscule, un chiffre, et un caractère spécial
func ValidatePassword(password string) []error {
    var errs []error

    if len(password) < 8 {
        errs = append(errs, Errors.ErrPasswordTooShort)
    }
    if !containsUppercase(password) {
        errs = append(errs, Errors.ErrPasswordNoUppercase)
    }
    if !containsLowercase(password) {
        errs = append(errs, Errors.ErrPasswordNoLowercase)
    }
    if !containsDigit(password) {
        errs = append(errs, Errors.ErrPasswordNoDigit)
    }
    if !containsSpecialChar(password) {
        errs = append(errs, Errors.ErrPasswordNoSpecialChar)
    }

    return errs
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