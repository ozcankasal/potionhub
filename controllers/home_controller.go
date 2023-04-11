package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (ctrl *HomeController) GetHome(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")

	var username string
	if user != nil {
		username = user.(string)
	}

	c.HTML(200, "home.html", gin.H{
		"title":    "Home",
		"username": username,
	})
}
