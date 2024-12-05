// cmd/Core/Errors/errors.go
package Errors

import "errors"

var (
    ErrUserNotFound             = errors.New("utilisateur non trouvé")
    ErrInvalidPassword          = errors.New("mot de passe incorrect")
    ErrValidationFailed         = errors.New("validation des champs échouée")
    ErrInternalServerError      = errors.New("erreur interne du serveur")
    ErrInvalidTokenClaims       = errors.New("les revendications du token sont invalides")
    ErrTokenInvalid             = errors.New("token invalide ou expiré. Veuillez vous authentifier")
    ErrTokenGeneration          = errors.New("erreur lors de la génération du token")
    ErrRefreshTokenGeneration   = errors.New("erreur lors de la génération du refresh token")
    ErrRefreshTokenInvalid      = errors.New("refresh token invalide ou expiré")
    ErrRefreshTokenMissing      = errors.New("refresh token manquant")
    ErrAccessTokenMissing       = errors.New("access token manquant")
    ErrEmailAlreadyUsed         = errors.New("cette adresse email est déjà utilisée")
    ErrPasswordsDoNotMatch      = errors.New("les mots de passe ne correspondent pas")
    ErrUsernameInvalidLength    = errors.New("le nom d'utilisateur doit contenir entre 3 et 20 caractères")
    ErrEmailInvalidFormat       = errors.New("l'adresse email n'est pas valide")
    ErrPasswordTooShort         = errors.New("le mot de passe doit contenir au moins 8 caractères")
    ErrPasswordNoUppercase      = errors.New("le mot de passe doit contenir au moins une majuscule")
    ErrPasswordNoLowercase      = errors.New("le mot de passe doit contenir au moins une minuscule")
    ErrPasswordNoDigit          = errors.New("le mot de passe doit contenir au moins un chiffre")
    ErrPasswordNoSpecialChar    = errors.New("le mot de passe doit contenir au moins un caractère spécial")
    // Ajoutez d'autres erreurs si nécessaire
)