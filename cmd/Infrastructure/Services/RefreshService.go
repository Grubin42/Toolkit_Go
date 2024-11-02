// cmd/Infrastructure/Services/RefreshService.go
package Services

import (
    "database/sql"
    "errors"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Models"
    "time"
)

type RefreshService struct {
    db *sql.DB
}

func NewRefreshService(db *sql.DB) *RefreshService {
    return &RefreshService{db: db}
}

// SaveRefreshToken enregistre un refresh token dans la base de données
func (s *RefreshService) SaveRefreshToken(userID int, token string, expiresAt time.Time) error {
    rt := &Models.RefreshToken{
        UserID:    userID,
        Token:     token,
        ExpiresAt: expiresAt,
    }
    return rt.Save(s.db)
}

// ValidateRefreshToken vérifie la validité d'un refresh token
func (s *RefreshService) ValidateRefreshToken(token string) (int, error) {
    rt := &Models.RefreshToken{}
    err := rt.FindByToken(s.db, token)
    if err != nil {
        return 0, err
    }

    if rt.Revoked {
        return 0, errors.New("refresh token révoqué")
    }

    if time.Now().After(rt.ExpiresAt) {
        return 0, errors.New("refresh token expiré")
    }

    return rt.UserID, nil
}

// RevokeRefreshToken révoque un refresh token spécifique
func (s *RefreshService) RevokeRefreshToken(token string) error {
    rt := &Models.RefreshToken{}
    err := rt.FindByToken(s.db, token)
    if err != nil {
        return err
    }

    return rt.Revoke(s.db)
}

// RevokeAllRefreshTokens révoque tous les refresh tokens d'un utilisateur
func (s *RefreshService) RevokeAllRefreshTokens(userID int) error {
    return Models.RevokeAllByUser(s.db, userID)
}