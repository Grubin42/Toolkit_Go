package Services

import (
    "database/sql"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Models"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
    "log"
    "net/http"
    "strings"
    "errors"
)

type RegisterService struct {
    db *sql.DB
}

func NewRegisterService(db *sql.DB) *RegisterService {
    return &RegisterService{db: db}
}

// RegisterUser enregistre un utilisateur après validation et hachage du mot de passe
func (rs *RegisterService) RegisterUser(username, email, password, confirmPassword string) (int, error) {
    // Validation du nom d'utilisateur
    if err := Utils.ValidateUsername(username); err != nil {
        return http.StatusBadRequest, err
    }

    // Validation de l'email
    if err := Utils.ValidateEmail(email); err != nil {
        return http.StatusBadRequest, err
    }

    // Validation du mot de passe
    passwordErrors := Utils.ValidatePassword(password)
    if len(passwordErrors) > 0 {
        return http.StatusBadRequest, errors.New(strings.Join(passwordErrors, ", "))
    }

    // Vérifier si les mots de passe correspondent
    if password != confirmPassword {
        return http.StatusBadRequest, errors.New("les mots de passe ne correspondent pas")
    }

    // Hachage du mot de passe
    passwordHash, err := Utils.HashPassword(password)
    if err != nil {
        log.Printf("Erreur lors du hachage du mot de passe : %v", err)
        return http.StatusInternalServerError, err
    }

    // Création d'un utilisateur
    user := Models.User{
        Name:         username,
        Email:        email,
        PasswordHash: passwordHash,
    }

    // Sauvegarde dans la base de données
    err = user.Save(rs.db)
    if err != nil {
        // Retourner une erreur HTTP spécifique si l'email est déjà utilisé
        if strings.Contains(err.Error(), "cette adresse email est déjà utilisée") {
            return http.StatusConflict, err // HTTP 409 Conflict
        }
        return http.StatusInternalServerError, err
    }

    // Retourner 201 Created en cas de succès
    return http.StatusCreated, nil
}