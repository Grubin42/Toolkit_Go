package Utils

import (
<<<<<<< HEAD:cmd/Infrastructure/Utils/auth.go
	"errors"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt/v5"
    "strings"
=======
    "net/http"
    "strings"
    "github.com/golang-jwt/jwt/v5"
    "time"
    "errors"
    "os"
>>>>>>> origin/gael-dev:cmd/Infrastructure/Utils/token.go
)

// Clé secrète utilisée pour signer les tokens (doit être sécurisée et stockée dans les variables d'environnement en production)
//var tokenExpiration = 24 * time.Hour 
var tokenExpiration = 20 * time.Second 

// GetTokenExpiration renvoie la durée d'expiration pour les contrôleurs
func GetTokenExpiration() time.Duration {
    return tokenExpiration
}

// jwtSecret retourne la clé secrète pour les JWT en la chargeant depuis les variables d'environnement
func getJWTSecret() []byte {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        panic("JWT_SECRET non défini dans les variables d'environnement") // Panique si la clé n'est pas définie
    }
    return []byte(secret)
}

// GenerateJWT génère un token JWT pour un utilisateur avec une expiration de 24 heures
func GenerateJWT(userID int) (string, error) {
    // Créer les claims du token
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(tokenExpiration).Unix(), // Expiration dans 24 heures
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
<<<<<<< HEAD:cmd/Infrastructure/Utils/auth.go
    // Signer le token avec la clé secrète
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
=======
    return token.SignedString(getJWTSecret())
>>>>>>> origin/gael-dev:cmd/Infrastructure/Utils/token.go
}

// ValidateJWT vérifie et valide un token JWT
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
    claims := jwt.MapClaims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return getJWTSecret(), nil
    })
    if err != nil || !token.Valid {
        return nil, errors.New("token invalide ou expiré. Veuillez-vous authentifier")
    }
    return claims, nil
}

<<<<<<< HEAD:cmd/Infrastructure/Utils/auth.go
func IsAuthentificated(r *http.Request) bool {
=======
// IsAuthenticated vérifie si l'utilisateur est authentifié
func IsAuthenticated(r *http.Request) bool {
>>>>>>> origin/gael-dev:cmd/Infrastructure/Utils/token.go
    cookie, err := r.Cookie("jwt_token")
    if err != nil {
        return false
    }

    token := cookie.Value
<<<<<<< HEAD:cmd/Infrastructure/Utils/auth.go
    return strings.TrimSpace(token) != ""
=======
    if strings.TrimSpace(token) == "" {
        return false
    }

    _, err = ValidateJWT(token)
    return err == nil
>>>>>>> origin/gael-dev:cmd/Infrastructure/Utils/token.go
}