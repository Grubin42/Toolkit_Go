// cmd/Core/Models/RefreshToken.go
package Models

import (
    "database/sql"
    "errors"
    "log"
    "strings"
    "time"
    "github.com/go-sql-driver/mysql"
)

type RefreshToken struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    Token     string    `json:"token"`
    ExpiresAt time.Time `json:"expires_at"`
    Revoked   bool      `json:"revoked"`
    CreatedAt time.Time `json:"created_at"`
}

// Save enregistre le refresh token dans la base de données
func (rt *RefreshToken) Save(db *sql.DB) error {
    if rt.UserID == 0 || rt.Token == "" || rt.ExpiresAt.IsZero() {
        return errors.New("les informations du refresh token sont incomplètes")
    }

    query := "INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES (?, ?, ?)"
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Printf("Erreur lors de la préparation de la requête d'insertion du refresh token : %v", err)
        return err
    }
    defer stmt.Close()

    res, err := stmt.Exec(rt.UserID, rt.Token, rt.ExpiresAt)
    if err != nil {
        // Vérification de l'erreur MySQL spécifique pour les duplicata
        if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
            if strings.Contains(mysqlErr.Message, "refresh_tokens.token") {
                return errors.New("ce refresh token existe déjà")
            }
        }
        log.Printf("Erreur lors de l'insertion du refresh token : %v", err)
        return err
    }

    id, err := res.LastInsertId()
    if err != nil {
        log.Printf("Erreur lors de la récupération de l'ID du refresh token inséré : %v", err)
        return err
    }
    rt.ID = int(id)

    return nil
}

// FindByToken récupère un refresh token à partir de son token
func (rt *RefreshToken) FindByToken(db *sql.DB, token string) error {
    query := "SELECT id, user_id, token, expires_at, revoked, created_at FROM refresh_tokens WHERE token = ? LIMIT 1"
    row := db.QueryRow(query, token)

    err := row.Scan(&rt.ID, &rt.UserID, &rt.Token, &rt.ExpiresAt, &rt.Revoked, &rt.CreatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return errors.New("refresh token non trouvé")
        }
        return err
    }

    return nil
}

// Revoke révoque ce refresh token spécifique
func (rt *RefreshToken) Revoke(db *sql.DB) error {
    query := "UPDATE refresh_tokens SET revoked = TRUE WHERE id = ?"
    res, err := db.Exec(query, rt.ID)
    if err != nil {
        log.Printf("Erreur lors de la révocation du refresh token : %v", err)
        return err
    }

    rowsAffected, err := res.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return errors.New("aucun refresh token révoqué")
    }

    rt.Revoked = true
    return nil
}

// RevokeAllByUser révoque tous les refresh tokens d'un utilisateur
func RevokeAllByUser(db *sql.DB, userID int) error {
    query := "UPDATE refresh_tokens SET revoked = TRUE WHERE user_id = ? AND revoked = FALSE"
    res, err := db.Exec(query, userID)
    if err != nil {
        log.Printf("Erreur lors de la révocation de tous les refresh tokens pour l'utilisateur %d : %v", userID, err)
        return err
    }

    rowsAffected, err := res.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        log.Printf("Aucun refresh token à révoquer pour l'utilisateur ID %d", userID)
        // Ne pas retourner d'erreur si aucun token n'a été révoqué
        return nil
    }

    log.Printf("Révocation de %d refresh tokens pour l'utilisateur ID %d", rowsAffected, userID)
    return nil
}