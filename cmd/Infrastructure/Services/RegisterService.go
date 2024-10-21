package Services

import (
	"github.com/Grubin42/Toolkit_Go/cmd/Core/Models"
	"errors"
	"log"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	db *sql.DB
}

func NewRegisterService(db *sql.DB) *RegisterService {
    return &RegisterService{
		db: db,
	}
}

func (rs *RegisterService) RegisterUser(username, email, password, confirmPassword string) error {
    // Validation des données
    if password != confirmPassword {
        return errors.New("les mots de passe ne correspondent pas")
    }
    
    if username == "" || email == "" || password == "" {
        return errors.New("tous les champs sont requis")
    }
    
    // Hash du mot de passe (tu devras utiliser un package de hashage comme bcrypt)
    passwordHash, erro := hashPassword(password)
	if erro != nil {
		log.Printf("Erreur lors du hachage du mot de passe : %v", erro)
		return erro // Retourner l'erreur si le hachage échoue
	}
    // Création de l'utilisateur
    user := Models.User{
        Name:     username,
        Email:        email,
        PasswordHash: passwordHash,
    }

    // Sauvegarde dans la base de données (cette partie dépend de ta configuration de la base de données)
    err := user.Save(rs.db)
    if err != nil {
        log.Printf("Erreur lors de l'enregistrement de l'utilisateur : %v", err)
        return err
    }

    return nil
}

func hashPassword(password string) (string, error) {
    // Générer le hachage avec un coût par défaut (bcrypt.DefaultCost)
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Erreur lors du hachage du mot de passe : %v", err)
        return "", err
    }

    // Retourner le mot de passe haché sous forme de chaîne
    return string(hashedPassword), nil
}