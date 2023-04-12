// controllers/welcome_controller.go

package controllers

import (
	"github.com/gin-gonic/gin"
)

type WelcomeController struct{}

func NewWelcomeController() *WelcomeController {
	return &WelcomeController{}
}

func (wc *WelcomeController) GetWelcome(c *gin.Context) {
	c.HTML(200, "welcome.html", gin.H{
		"title": "Welcome",
	})
}
