package controllers

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (ctrl *HomeController) GetHome(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")

	fmt.Println("Username in session:", username)

	c.HTML(200, "home.html", gin.H{
		"title":    "Home",
		"username": username.(string),
	})
}
