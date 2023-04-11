package main

import (
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ozcankasal/potionhub/controllers"
)

func main() {
	r := gin.Default()

	// Register custom template functions
	funcMap := template.FuncMap{
		"year": func() int {
			return time.Now().Year()
		},
	}

	// Load templates with custom functions
	r.SetFuncMap(funcMap)
	r.LoadHTMLGlob("templates/*")

	homeController := controllers.NewHomeController()
	r.GET("/", homeController.GetHome)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
