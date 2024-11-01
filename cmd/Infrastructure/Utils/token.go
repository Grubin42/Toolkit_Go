package Utils

import (
    "net/http"
    "strings"
    "github.com/golang-jwt/jwt/v5"
    "time"
    "errors"
    "os"
)

// Durées de validité des tokens
var (
    AccessTokenExpiration  = 1 * time.Hour
    RefreshTokenExpiration = 30 * 24 * time.Hour // 30 jours
)

// GetAccessTokenExpiration renvoie la durée d'expiration pour les access tokens
func GetAccessTokenExpiration() time.Duration {
    return AccessTokenExpiration
}

// GetRefreshTokenExpiration renvoie la durée d'expiration pour les refresh tokens
func GetRefreshTokenExpiration() time.Duration {
    return RefreshTokenExpiration
}

// jwtSecret retourne la clé secrète pour les JWT en la chargeant depuis les variables d'environnement
func getJWTSecret() []byte {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        panic("JWT_SECRET non défini dans les variables d'environnement")
    }
    return []byte(secret)
}

// GenerateAccessToken génère un access token JWT pour un utilisateur avec une expiration de 1 heure
func GenerateAccessToken(userID int) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(AccessTokenExpiration).Unix(),
        "type":    "access",
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(getJWTSecret())
}

// GenerateRefreshToken génère un refresh token JWT pour un utilisateur avec une expiration de 30 jours
func GenerateRefreshToken(userID int) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(RefreshTokenExpiration).Unix(),
        "type":    "refresh",
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(getJWTSecret())
}

// ValidateJWT vérifie et valide un token JWT
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
    claims := jwt.MapClaims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return getJWTSecret(), nil
    })
    if err != nil || !token.Valid {
        return nil, errors.New("token invalide ou expiré")
    }
    return claims, nil
}

// IsAuthenticated vérifie si l'utilisateur est authentifié
func IsAuthenticated(r *http.Request) bool {
    cookie, err := r.Cookie("jwt_token")
    if err != nil {
        return false
    }

    token := cookie.Value
    if strings.TrimSpace(token) == "" {
        return false
    }

    _, err = ValidateJWT(token)
    return err == nil
}