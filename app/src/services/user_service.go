package services

import (
	"database/sql"
	//"log"
	"toolkit_go/src/models"
)

func GetAllUsers(db *sql.DB) ([]models.User, error)  {
	var users []models.User
	
	rows, err := db.Query("SELECT id, username, email FROM users")
	if err != nil {
        return nil, err
    }
	defer rows.Close()
	for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
	if err := rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}