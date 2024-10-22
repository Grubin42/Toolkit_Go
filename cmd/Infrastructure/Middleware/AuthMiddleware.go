package Middleware

import (
    "context"
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
)

type key string

const userIDKey key = "userID"

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Récupérer le cookie contenant le token
        cookie, err := r.Cookie("jwt_token")
        if err != nil {
            http.Error(w, "Non autorisé", http.StatusUnauthorized)
            return
        }

        // Valider le token et extraire les claims
        claims, err := Utils.ValidateJWT(cookie.Value)
        if err != nil {
            http.Error(w, "Non autorisé : "+err.Error(), http.StatusUnauthorized)
            return
        }

        // Extraire l'ID de l'utilisateur depuis les claims
        userID, ok := claims["user_id"].(float64) // JWT stocke les nombres en float64
        if !ok {
            http.Error(w, "Token invalide", http.StatusUnauthorized)
            return
        }

        // Ajouter l'ID de l'utilisateur dans le contexte de la requête
        ctx := context.WithValue(r.Context(), userIDKey, int(userID))
        r = r.WithContext(ctx)

        // Passer à l'étape suivante avec le contexte mis à jour
        next.ServeHTTP(w, r)
    })
}