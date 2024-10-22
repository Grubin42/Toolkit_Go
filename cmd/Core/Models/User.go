package Models

import (
    "database/sql"
    "errors"
    "log"
    "strings"
    "github.com/go-sql-driver/mysql" // Import du driver MySQL pour gérer les erreurs spécifiques
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    PasswordHash string `json:"PasswordHash"`
}

// Save enregistre l'utilisateur dans la base de données en utilisant des requêtes SQL natives
func (u *User) Save(db *sql.DB) error {
    if u.Name == "" || u.Email == "" || u.PasswordHash == "" {
        return errors.New("les informations de l'utilisateur sont incomplètes")
    }

    // Préparer l'instruction d'insertion
    query := "INSERT INTO users (name, email, passwordhash) VALUES (?, ?, ?)"
    stmt, err := db.Prepare(query)
    if err != nil {
        log.Printf("Erreur lors de la préparation de la requête : %v", err)
        return err
    }
    defer stmt.Close()

    // Exécuter l'insertion
    _, err = stmt.Exec(u.Name, u.Email, u.PasswordHash)
    if err != nil {
        // Vérification de l'erreur MySQL spécifique pour les duplicata
        if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
            // Intercepter l'erreur de duplicata et retourner un message clair
            if strings.Contains(mysqlErr.Message, "users.email") {
                return errors.New("cette adresse email est déjà utilisée")
            }
        }
        log.Printf("Erreur lors de l'insertion de l'utilisateur : %v", err)
        return err
    }

    return nil
}

// FindByUsernameOrEmail récupère un utilisateur à partir de son nom d'utilisateur ou de son email
func (u *User) FindByUsernameOrEmail(db *sql.DB, identifier string) error {
    query := "SELECT id, name, email, passwordhash FROM users WHERE name = ? OR email = ? LIMIT 1"
    row := db.QueryRow(query, identifier, identifier)
    
    err := row.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash)
    if err != nil {
        if err == sql.ErrNoRows {
            return errors.New("utilisateur non trouvé")
        }
        return err
    }

    return nil
}