// controllers/home_controller.go

package controllers

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// HomeController is the wise and welcoming spirit that greets travelers
// upon their arrival to the realm of PotionHub. It provides guidance
// and comfort, ensuring that they feel at home in this mystical land.
type HomeController struct{}

// NewHomeController invokes the kind-hearted HomeController spirit
// to guide and assist the travelers in their journey.
func NewHomeController() *HomeController {
	return &HomeController{}
}

// GetHome is the HomeController's warm embrace, a place where travelers
// can rest and gather their strength before venturing forth into the
// magical world of PotionHub.
func (ctrl *HomeController) GetHome(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")

	fmt.Println("Username in session:", username)

	c.HTML(200, "home.html", gin.H{
		"title":    "Home",
		"username": username.(string),
	})
}
