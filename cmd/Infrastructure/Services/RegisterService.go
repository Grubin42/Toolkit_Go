package Services

import (
    "database/sql"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Models"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Errors"
)

type RegisterService struct {
    db *sql.DB
}

func NewRegisterService(db *sql.DB) *RegisterService {
    return &RegisterService{db: db}
}

// RegisterUser enregistre un utilisateur après validation et hachage du mot de passe
func (rs *RegisterService) RegisterUser(username, email, password, confirmPassword string) ([]error, error) {
    var validationErrors []error

    // Validation du nom d'utilisateur
    if err := Utils.ValidateUsername(username); err != nil {
        validationErrors = append(validationErrors, err)
    }

    // Validation de l'email
    if err := Utils.ValidateEmail(email); err != nil {
        validationErrors = append(validationErrors, err)
    }

    // Validation du mot de passe
    passwordErrors := Utils.ValidatePassword(password)
    if len(passwordErrors) > 0 {
        validationErrors = append(validationErrors, passwordErrors...)
    }

    // Vérifier si les mots de passe correspondent
    if password != confirmPassword {
        validationErrors = append(validationErrors, Errors.ErrPasswordsDoNotMatch)
    }
    
    if len(validationErrors) > 0 {
        // Retourner une erreur de validation personnalisée
        return validationErrors, nil
    }
    // Hachage du mot de passe
    passwordHash, err := Utils.HashPassword(password)
    if err != nil {
        return nil, Errors.ErrInternalServerError
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
        // Retourner une erreur spécifique si l'email est déjà utilisé
        if err == Errors.ErrEmailAlreadyUsed {
            return nil, err
        }
        return nil, Errors.ErrInternalServerError
    }

    // Retourner 201 Created en cas de succès
    return nil, nil
}