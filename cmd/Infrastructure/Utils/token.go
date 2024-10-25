package Utils

import (
    "github.com/golang-jwt/jwt/v5"
    "time"
    "errors"
)

// Clé secrète utilisée pour signer les tokens (doit être sécurisée et stockée dans les variables d'environnement en production)
var jwtSecret = []byte("my_secret_key")

// GenerateJWT génère un token JWT pour un utilisateur avec une expiration de 24 heures
func GenerateJWT(userID int) (string, error) {
    // Créer les claims du token
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(), // Expiration dans 24 heures
    }

    // Créer un nouveau token avec la méthode de signature HMAC et les claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Signer le token avec la clé secrète
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// ValidateJWT vérifie et valide un token JWT, et retourne les claims s'il est valide
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
    claims := jwt.MapClaims{}

    // Parser et valider le token
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    // Vérifier si le token est valide
    if err != nil || !token.Valid {
        return nil, errors.New("token invalide ou expiré")
    }

    return claims, nil
}