package Models

// FormErrors structure pour gérer les erreurs de formulaire
type FormErrors struct {
    ErrorMessage string            `json:"error_message"`
    FieldErrors  map[string]string `json:"field_errors"`
}