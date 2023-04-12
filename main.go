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

	// Conjure custom template functions for our magical templates
	funcMap := template.FuncMap{
		// Reveal the current year to the mystical beings of PotionHub
		"year": func() int {
			return time.Now().Year()
		},
		// Transform potion prices into dazzling dollar format
		"formatAsDollars": utils.FormatAsDollars,
		// Join strings as if weaving a powerful incantation
		"join": strings.Join,
	}

	// Enchant the Gin router with custom functions and summon the templates
	r.SetFuncMap(funcMap)
	r.LoadHTMLGlob("templates/*.html")

	// Set up a magical session store to remember the identity of visitors
	store := cookie.NewStore([]byte("secret-key"))
	r.Use(sessions.Sessions("mysession", store))

	// Prepare the LoggedIn middleware to guard our enchanted routes
	loggedInMiddleware := middleware.LoggedIn()

	// Unveil the catalog of potions, protected by the LoggedIn middleware
	catalogController := controllers.NewCatalogController()
	r.GET("/catalog", loggedInMiddleware, catalogController.GetCatalog)

	// Reveal the entrance to the authentication chamber
	authController := controllers.NewAuthController()
	r.GET("/login", authController.GetLogin)

	// Secure the authentication chamber with a mystical password lock
	authMiddleware := middleware.AuthRequired("SpellBinder", "ElixirOfEnigma")
	r.POST("/login", authMiddleware, authController.PostLogin)

	// Greet newcomers to PotionHub with a warm and welcoming spell
	welcomeController := controllers.NewWelcomeController()
	r.GET("/welcome", welcomeController.GetWelcome)

	// Display the magical home of PotionHub, accessible only to worthy adventurers
	homeController := controllers.NewHomeController()
	r.GET("/", loggedInMiddleware, homeController.GetHome)

	// Create a secret passage for adventurers to leave the mystical realm
	r.GET("/logout", authController.Logout)

	// Summon the Gin framework to listen for and serve the magical beings
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
