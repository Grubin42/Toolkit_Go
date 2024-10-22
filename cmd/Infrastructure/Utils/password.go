package Utils

import (
    "golang.org/x/crypto/bcrypt"
    "log"
)

// HashPassword génère un hachage sécurisé pour un mot de passe donné
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Erreur lors du hachage du mot de passe : %v", err)
        return "", err
    }
    return string(hashedPassword), nil
}

// CheckPassword compare un mot de passe fourni avec un hachage stocké
func CheckPassword(hashedPassword, password string) error {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        return err
    }
    return nil
}