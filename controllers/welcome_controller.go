// controllers/welcome_controller.go

package controllers

import (
	"github.com/gin-gonic/gin"
)

// WelcomeController is the mystical gatekeeper of the PotionHub realm,
// standing tall at the entrance to this enchanted land. It warmly
// greets newcomers, inviting them to embark on a journey filled with
// wonder, magic, and discovery.
type WelcomeController struct{}

// NewWelcomeController summons the noble WelcomeController to stand
// guard at the gates of PotionHub, ensuring that all who enter are
// welcomed with open arms.
func NewWelcomeController() *WelcomeController {
	return &WelcomeController{}
}

// GetWelcome is the warm and inviting embrace of the WelcomeController,
// beckoning travelers to step into the enchanted world of PotionHub.
// With a wave of its hand, the WelcomeController unveils the magical
// realm before the traveler's eyes, filling their hearts with excitement
// and anticipation.
func (wc *WelcomeController) GetWelcome(c *gin.Context) {
	c.HTML(200, "welcome.html", gin.H{
		"title": "Welcome",
	})
}
