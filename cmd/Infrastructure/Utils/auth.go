package Utils

import (
    "net/http"
    "strings"
    "github.com/golang-jwt/jwt/v5"
    "time"
    "errors"
    "os"
    "encoding/json"
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

func SetAccessTokenCookie(w http.ResponseWriter, token string) {
    http.SetCookie(w, &http.Cookie{
        Name:     "jwt_token",
        Value:    token,
        Expires:  time.Now().Add(GetAccessTokenExpiration()),
        HttpOnly: true,
        Secure:   true, // Activer en production (HTTPS)
        SameSite: http.SameSiteStrictMode,
        Path:     "/",
    })
}

func SetRefreshTokenCookie(w http.ResponseWriter, token string) {
    http.SetCookie(w, &http.Cookie{
        Name:     "refresh_token",
        Value:    token,
        Expires:  time.Now().Add(GetRefreshTokenExpiration()),
        HttpOnly: true,
        Secure:   true, // Activer en production (HTTPS)
        SameSite: http.SameSiteStrictMode,
        Path:     "/refresh",
    })
}

func ClearTokenCookies(w http.ResponseWriter) {
    http.SetCookie(w, &http.Cookie{
        Name:     "jwt_token",
        Value:    "",
        Expires:  time.Unix(0, 0),
        HttpOnly: true,
        Secure:   true, // Activer en production (HTTPS)
        SameSite: http.SameSiteStrictMode,
        Path:     "/",
    })

    http.SetCookie(w, &http.Cookie{
        Name:     "refresh_token",
        Value:    "",
        Expires:  time.Unix(0, 0),
        HttpOnly: true,
        Secure:   true, // Activer en production (HTTPS)
        SameSite: http.SameSiteStrictMode,
        Path:     "/refresh",
    })
}

// getJWTSecret retourne la clé secrète pour les JWT en la chargeant depuis les variables d'environnement
func getJWTSecret() ([]byte, error) {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        return nil, errors.New("JWT_SECRET non défini dans les variables d'environnement")
    }
    return []byte(secret), nil
}

// GenerateAccessToken génère un access token JWT pour un utilisateur avec une expiration de 1 heure
func GenerateAccessToken(userID int) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(AccessTokenExpiration).Unix(),
        "type":    "access",
    }

    secret, err := getJWTSecret()
    if err != nil {
        return "", err
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secret)
}

// GenerateRefreshToken génère un refresh token JWT pour un utilisateur avec une expiration de 30 jours
func GenerateRefreshToken(userID int) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(RefreshTokenExpiration).Unix(),
        "type":    "refresh",
    }

    secret, err := getJWTSecret()
    if err != nil {
        return "", err
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secret)
}

// GenerateTokens génère à la fois l'access token et le refresh token
func GenerateTokens(userID int) (accessToken string, refreshToken string, err error) {
    accessToken, err = GenerateAccessToken(userID)
    if err != nil {
        return
    }

    refreshToken, err = GenerateRefreshToken(userID)
    return
}

// ValidateJWT vérifie et valide un token JWT
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
    claims := jwt.MapClaims{}

    secret, err := getJWTSecret()
    if err != nil {
        return nil, err
    }

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return secret, nil
    })
    if err != nil || !token.Valid {
        return nil, errors.New("token invalide ou expiré. Veuillez-vous authentifier")
    }

    return claims, nil
}

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

// ErrorResponse struct pour les réponses d'erreur
type ErrorResponse struct {
    Message string            `json:"message"`
    Details map[string]string `json:"details,omitempty"`
}

// WriteError écrit une réponse d'erreur JSON
func WriteError(w http.ResponseWriter, status int, message string, details map[string]string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(ErrorResponse{
        Message: message,
        Details: details,
    })
}