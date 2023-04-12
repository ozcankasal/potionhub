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
	"github.com/ozcankasal/potionhub/models"
	"github.com/ozcankasal/potionhub/services"
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

	// Initialize the database connection
	db, err := models.NewDB("postgres://postgres:admin@0.0.0.0/potion_catalog?sslmode=disable")
	if err != nil {
		panic(err)
	}

	userService := services.NewUserService(db)

	// Instantiate the controllers with the UserService
	catalogController := controllers.NewCatalogController()
	authController := controllers.NewAuthController(userService)

	// Set up routes
	r.GET("/catalog", catalogController.GetCatalog)
	r.GET("/login", authController.GetLogin)

	// Create and set up the AuthRequired middleware
	authMiddleware := middleware.AuthRequired(userService)
	r.POST("/login", authMiddleware, authController.PostLogin)

	welcomeController := controllers.NewWelcomeController()
	r.GET("/welcome", welcomeController.GetWelcome)

	registerController := controllers.NewRegisterController(userService)
	// Add the registration routes
	r.GET("/register", registerController.GetRegister)
	r.POST("/register", registerController.PostRegister)

	homeController := controllers.NewHomeController()
	loggedInMiddleware := middleware.LoggedIn()
	r.GET("/", loggedInMiddleware, homeController.GetHome)

	// Logout route
	r.GET("/logout", authController.Logout)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
