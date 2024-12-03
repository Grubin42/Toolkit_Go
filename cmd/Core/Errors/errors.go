package Errors

const (
    ErrorLoginFailed              = "Échec de la connexion. Veuillez vérifier vos identifiants."
    ErrorTokenGeneration          = "Erreur lors de la génération du token."
    ErrorRefreshTokenGeneration   = "Erreur lors de la génération du refresh token."
    ErrorRefreshTokenSave         = "Erreur lors de l'enregistrement du refresh token."
    ErrorRevokeAllTokens          = "Erreur lors de la réinitialisation des tokens existants."
)
// ErrorType représente le type d'erreur
type ErrorType string

const (
    ValidationError     ErrorType = "ValidationError"
    AuthenticationError ErrorType = "AuthenticationError"
    ServerError         ErrorType = "ServerError"
)

// AppError structure pour gérer les erreurs avec un type
type AppError struct {
    Type    ErrorType
    Message string
    Field   string // Utilisé pour les erreurs de validation
}

func (e *AppError) Error() string {
    return e.Message
}

// Fonctions utilitaires pour créer des erreurs spécifiques
func NewValidationError(field, message string) *AppError {
    return &AppError{
        Type:    ValidationError,
        Message: message,
        Field:   field,
    }
}

func NewAuthenticationError(message string) *AppError {
    return &AppError{
        Type:    AuthenticationError,
        Message: message,
    }
}

func NewServerError(message string) *AppError {
    return &AppError{
        Type:    ServerError,
        Message: message,
    }
}