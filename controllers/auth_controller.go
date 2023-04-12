// controllers/auth_controller.go

package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthController is the sentinel of the enchanted portal to the realm of PotionHub.
// It guards the passage against unwelcome visitors, ensuring that only those with
// the correct credentials may enter.
type AuthController struct{}

// NewAuthController summons the AuthController to begin its watch over the sacred portal.
func NewAuthController() *AuthController {
	return &AuthController{}
}

// GetLogin presents the mystical login parchment that travelers must complete
// to request passage into the realm.
func (ctrl *AuthController) GetLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

// PostLogin examines the traveler's completed parchment, seeking the true essence
// of their identity hidden within the ancient runes.
func (ctrl *AuthController) PostLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// The sacred runes that reveal one's true essence
	mockUsername := "SpellBinder"
	mockPassword := "ElixirOfEnigma"

	session := sessions.Default(c)

	if username == mockUsername && password == mockPassword {
		// Engrave the traveler's essence into the spirit of the realm
		session.Set("user", username)
		_ = session.Save()

		// Grant passage through the enchanted portal
		c.Redirect(http.StatusSeeOther, "/")
	} else {
		// Deny the traveler's request, presenting them with the parchment once more
		c.HTML(200, "login.html", gin.H{
			"error": "Invalid username or password",
		})
	}
}

// Logout is a powerful incantation that severs the connection between a traveler
// and the realm, casting them back to the mundane world outside.
func (ac *AuthController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	_ = session.Save()

	c.Redirect(http.StatusSeeOther, "/")
}
