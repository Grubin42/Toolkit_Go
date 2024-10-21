package Models

import (
    "database/sql"
    "errors"
    "log"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    PasswordHash string `json:"PasswordHash"`
}

// Exemple de fonction de sauvegarde (à remplacer par la logique de ton repository)
// Save enregistre l'utilisateur dans la base de données en utilisant des requêtes SQL natives
func (u *User) Save(db *sql.DB) error {
    // Vérifier si les informations obligatoires sont présentes
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
    result, err := stmt.Exec(u.Name, u.Email, u.PasswordHash)
    if err != nil {
        log.Printf("Erreur lors de l'insertion de l'utilisateur : %v", err)
        return err
    }

    // Récupérer l'ID généré
    lastInsertID, err := result.LastInsertId()
    if err != nil {
        log.Printf("Erreur lors de la récupération de l'ID de l'utilisateur : %v", err)
        return err
    }

    u.ID = int(lastInsertID)
    return nil
}