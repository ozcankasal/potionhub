// auth_controller.go

package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ctrl *AuthController) GetLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func (ctrl *AuthController) PostLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Mock username and password for validation
	mockUsername := "admin"
	mockPassword := "admin"

	session := sessions.Default(c)

	if username == mockUsername && password == mockPassword {
		// Set the user session
		session.Set("user", username)
		_ = session.Save()

		// Redirect to the home page
		c.Redirect(http.StatusSeeOther, "/")
	} else {
		// Show an error message and render the login template again
		c.HTML(200, "login.html", gin.H{
			"error": "Invalid username or password",
		})
	}
}

func (ac *AuthController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	_ = session.Save()

	c.Redirect(http.StatusSeeOther, "/")
}
