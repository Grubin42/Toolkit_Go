package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql" // Import du driver MySQL
)

// ConnectDB établit la connexion à la base de données MySQL avec un mécanisme de retry
func ConnectDB() (*sql.DB, error) {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	// Format de la chaîne de connexion pour MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var db *sql.DB
	var err error

	// Retry mechanism: essayer de se connecter jusqu'à 10 fois, avec une attente de 2 secondes entre chaque tentative
	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil && db.Ping() == nil {
			return db, nil
		}

		log.Printf("Tentative de connexion à MySQL échouée. Réessai dans 2 secondes... (%d/10)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("impossible de se connecter à MySQL: %v", err)
}