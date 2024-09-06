package Services

import (
    "database/sql"
    "github.com/Grubin42/Toolkit_Go/internal/models"
)

type UserService struct {
    db *sql.DB
}

// NewUserService initialise un nouveau service utilisateur.
func NewUserService(db *sql.DB) *UserService {
    return &UserService{
        db: db,
    }
}

// CreateUser ajoute un nouvel utilisateur à la base de données.
func (us *UserService) CreateUser(user *models.User) error {
    query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
    return us.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
}

// GetUserByID récupère un utilisateur par son ID.
func (us *UserService) GetUserByID(id int) (*models.User, error) {
    var user models.User
    query := "SELECT id, name, email FROM users WHERE id = $1"
    err := us.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

// UpdateUser met à jour les informations d'un utilisateur.
func (us *UserService) UpdateUser(user *models.User) error {
    query := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
    _, err := us.db.Exec(query, user.Name, user.Email, user.ID)
    return err
}

// DeleteUser supprime un utilisateur par son ID.
func (us *UserService) DeleteUser(id int) error {
    query := "DELETE FROM users WHERE id = $1"
    _, err := us.db.Exec(query, id)
    return err
}

func (us *UserService) GetAllUsers() ([]models.User, error) {
    var users []models.User
    query := "SELECT id, name, email FROM users"
    rows, err := us.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}