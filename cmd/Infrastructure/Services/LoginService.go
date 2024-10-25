package Services

import (
    "database/sql"
    "errors"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Models"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
    "net/http"
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
func (ls *LoginService) LoginUser(identifier, password string) (int, string, error) {
    var user Models.User

    // Rechercher l'utilisateur par nom d'utilisateur ou email
    err := user.FindByUsernameOrEmail(ls.db, identifier)
    if err != nil {
        return http.StatusUnauthorized, "", errors.New("identifiant ou mot de passe incorrect")
    }

    // Vérifier le mot de passe
    err = Utils.CheckPassword(user.PasswordHash, password)
    if err != nil {
        return http.StatusUnauthorized, "", errors.New("identifiant ou mot de passe incorrect")
    }

    // Générer le token JWT en utilisant l'utilitaire
    token, err := Utils.GenerateJWT(user.ID)
    if err != nil {
        return http.StatusInternalServerError, "", err
    }

    return http.StatusOK, token, nil
}