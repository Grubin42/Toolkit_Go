package routes

import (
    "github.com/gin-gonic/gin"
    "toolkit_go/src/controllers"
)

func LoadRoutes(r *gin.Engine) {
    r.GET("/", controllers.HomePage)
}