package controllers

import (
    "database/sql"
    "encoding/json"
    "html/template"
    "net/http"
    "strconv"
	"strings"
    "github.com/Grubin42/Toolkit_Go/internal/models"
    "github.com/Grubin42/Toolkit_Go/internal/services"
)

type UserController struct {
    service   *services.UserService
    templates *template.Template
}

func NewUserController(db *sql.DB) *UserController {
    templates := template.Must(template.ParseFiles("templates/index.html", "templates/users.html"))
    return &UserController{
        service:   services.NewUserService(db),
        templates: templates,
    }
}

func (uc *UserController) HandleUsers(w http.ResponseWriter, r *http.Request) {
    id := strings.TrimPrefix(r.URL.Path, "/users/")
    if id == "" {
        if r.Method == http.MethodPost {
            uc.createUser(w, r)
        } else if r.Method == http.MethodGet {
            uc.ListUsers(w, r) // Assurez-vous que cette méthode est publique
        } else {
            w.WriteHeader(http.StatusMethodNotAllowed)
        }
    } else {
        switch r.Method {
        case http.MethodGet:
            uc.getUser(w, r, id)
        case http.MethodPut:
            uc.updateUser(w, r, id)
        case http.MethodDelete:
            uc.deleteUser(w, r, id)
        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
        }
    }
}

func (uc *UserController) HandleIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        uc.templates.ExecuteTemplate(w, "index.html", nil)
    } else {
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

// Renommez la méthode en ListUsers pour qu'elle soit exportée
func (uc *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
    users, err := uc.service.GetAllUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    uc.templates.ExecuteTemplate(w, "users.html", users)
}

func (uc *UserController) createUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = uc.service.CreateUser(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func (uc *UserController) getUser(w http.ResponseWriter, r *http.Request, id string) {
    userID, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    user, err := uc.service.GetUserByID(userID)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    json.NewEncoder(w).Encode(user)
}

func (uc *UserController) updateUser(w http.ResponseWriter, r *http.Request, id string) {
    userID, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var user models.User
    err = json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user.ID = userID
    err = uc.service.UpdateUser(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}

func (uc *UserController) deleteUser(w http.ResponseWriter, r *http.Request, id string) {
    userID, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    err = uc.service.DeleteUser(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}