package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Import du driver MySQL
	"os"
)

// ConnectDB établit la connexion à la base de données MySQL
func ConnectDB() (*sql.DB, error) {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	// Format de la chaîne de connexion pour MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Vérifier la connexion en effectuant un ping
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
