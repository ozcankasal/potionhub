package main

import (
	"html/template"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/ozcankasal/potionhub/controllers"
	"github.com/ozcankasal/potionhub/middleware"
	"github.com/ozcankasal/potionhub/utils"
)

func main() {
	r := gin.Default()

	// Register custom template functions
	funcMap := template.FuncMap{
		"year": func() int {
			return time.Now().Year()
		},
		"formatAsDollars": utils.FormatAsDollars,
		"join":            strings.Join,
	}

	// Load templates with custom functions
	r.SetFuncMap(funcMap)
	r.LoadHTMLGlob("templates/*.html")

	// Set up session middleware
	store := cookie.NewStore([]byte("secret-key"))
	r.Use(sessions.Sessions("mysession", store))

	loggedInMiddleware := middleware.LoggedIn()

	catalogController := controllers.NewCatalogController()
	r.GET("/catalog", loggedInMiddleware, catalogController.GetCatalog)

	authController := controllers.NewAuthController()
	r.GET("/login", authController.GetLogin)

	// Replace "fixedUsername" and "fixedPassword" with your desired values
	authMiddleware := middleware.AuthRequired("admin", "admin")
	r.POST("/login", authMiddleware, authController.PostLogin)

	welcomeController := controllers.NewWelcomeController()
	r.GET("/welcome", welcomeController.GetWelcome)

	homeController := controllers.NewHomeController()
	r.GET("/", loggedInMiddleware, homeController.GetHome)

	// Logout route
	r.GET("/logout", authController.Logout)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
