package Middleware

import (
    "time"
    "net/http"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("jwt_token")
        if err != nil {
            if err == http.ErrNoCookie {
                http.Redirect(w, r, "/login", http.StatusSeeOther)
                return
            }
            http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
            return
        }

        tokenString := cookie.Value
        if tokenString == "" {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        _, err = Utils.ValidateJWT(tokenString)
        if err != nil {
            http.SetCookie(w, &http.Cookie{
                Name:     "jwt_token",
                Value:    "",
                Expires:  time.Unix(0, 0),
                HttpOnly: true,
                Secure:   true,
                Path:     "/",
            })
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        next.ServeHTTP(w, r)
    })
}