package Controllers

import (
    "net/http"
    "database/sql"
    "errors"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Services"
    "github.com/Grubin42/Toolkit_Go/cmd/Infrastructure/Utils"
    "github.com/Grubin42/Toolkit_Go/cmd/Core/Errors"
)

type RegisterController struct {
    BaseController
    registerService *Services.RegisterService
}

func NewRegisterController(db *sql.DB) *RegisterController {
    return &RegisterController{
        BaseController{
            Templates: Utils.LoadTemplates(
                "Register/index.html",
            ),
        },
        Services.NewRegisterService(db),
    }
}

func (rc *RegisterController) HandleIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        if err := r.ParseForm(); err != nil {
            rc.HandleError(w, r, "Inscription", "Erreur lors de la soumission du formulaire.", nil, http.StatusBadRequest)
            return
        }

        username := r.FormValue("username")
        email := r.FormValue("email")
        password := r.FormValue("password")
        confirmPassword := r.FormValue("confirm_password")
        fieldErrors := make(map[string]string)

        validationErrors, err := rc.registerService.RegisterUser(username, email, password, confirmPassword)
        if len(validationErrors) > 0 || err != nil {
            if len(validationErrors) > 0 {
                // Parcourir les erreurs de validation pour remplir fieldErrors
                for _, ve := range validationErrors {
                    switch {
                    case errors.Is(ve, Errors.ErrUsernameInvalidLength):
                        fieldErrors["username"] = ve.Error()
                    case errors.Is(ve, Errors.ErrEmailInvalidFormat):
                        fieldErrors["email"] = ve.Error()
                    case errors.Is(ve, Errors.ErrPasswordsDoNotMatch):
                        fieldErrors["confirm_password"] = ve.Error()
                    case errors.Is(ve, Errors.ErrPasswordTooShort),
                         errors.Is(ve, Errors.ErrPasswordNoUppercase),
                         errors.Is(ve, Errors.ErrPasswordNoLowercase),
                         errors.Is(ve, Errors.ErrPasswordNoDigit),
                         errors.Is(ve, Errors.ErrPasswordNoSpecialChar):
                        // Concaténer les erreurs liées au mot de passe
                        if existing, ok := fieldErrors["password"]; ok {
                            fieldErrors["password"] = existing + ", " + ve.Error()
                        } else {
                            fieldErrors["password"] = ve.Error()
                        }
                    default:
                        // Autres erreurs de validation
                        fieldErrors["general"] = ve.Error()
                    }
                }
                rc.HandleError(w, r, "Inscription", "Validation des champs échouée.", fieldErrors, http.StatusBadRequest)
            } else if err != nil {
                // Gestion des autres erreurs
                switch {
                case errors.Is(err, Errors.ErrEmailAlreadyUsed):
                    fieldErrors["email"] = err.Error()
                    rc.HandleError(w, r, "Inscription", "Validation des champs échouée.", fieldErrors, http.StatusConflict)
                default:
                    rc.HandleError(w, r, "Inscription", "Une erreur est survenue.", nil, http.StatusInternalServerError)
                }
            }
            return
        }

        w.Header().Set("HX-Redirect", "/login")
        return
    }

    specificData := map[string]interface{}{
        "Title": "Inscription",
    }

    rc.Render(w, r, specificData)
}