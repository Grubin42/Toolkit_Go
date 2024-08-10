package controllers

import (
    "github.com/gin-gonic/gin"
    "toolkit_go/src/services"
    //"toolkit_go/models"
    "database/sql"
    "net/http"
    "log"
)

var db *sql.DB // La connexion à la base de données, à initialiser dans `main.go`

func HomePage(c *gin.Context) {
    users, err := services.GetAllUsers(db)
    if err != nil {
        log.Println("Error fetching users:", err)
        c.String(http.StatusInternalServerError, "Error fetching users")
        return
    }

    c.HTML(http.StatusOK, "home/index.html", gin.H{
        "title": "Home",
        "users": users,
    })
}