package Services

import (
    "database/sql"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Models"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Errors"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
)

type LoginService struct {
    db *sql.DB
}

func NewLoginService(db *sql.DB) *LoginService {
    return &LoginService{
        db: db,
    }
}

// LoginUser vérifie si l'utilisateur existe et si le mot de passe est correct, puis génère un JWT
func (ls *LoginService) LoginUser(identifier, password string) (bool, int, error) {
    var user Models.User
    
    // Rechercher l'utilisateur par nom d'utilisateur ou email
    err := user.FindByUsernameOrEmail(ls.db, identifier)
    if err != nil {
        // Retourne une erreur spécifique pour le champ username si l'utilisateur n'est pas trouvé
        return false, 0, Errors.NewAuthenticationError("username_not_found")
    }

    // Vérifier le mot de passe
    err = Utils.CheckPassword(user.PasswordHash, password)
    if err != nil {
        // Retourne une erreur spécifique pour le champ password si le mot de passe est incorrect
        return false, 0, Errors.NewAuthenticationError("incorrect_password")
    }

    return true, user.ID, nil
}