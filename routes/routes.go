package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

func SetupRoutes(r *gin.Engine) {

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":   "My Go application",
			"Heading": "Hello, world!",
			"Message": "Welcome to my Go web application with Gin and HTML templates",
		})
	})

	r.Static("/static", "./static")
}
