package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ozcankasal/potionhub/services"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	userService *services.UserService
}

func NewAuthController(userService *services.UserService) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

func (ctrl *AuthController) GetLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func (ctrl *AuthController) PostLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := ctrl.userService.GetUserByUsername(username)
	if err != nil {
		c.HTML(200, "login.html", gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Compare the provided password with the hashed password from the database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.HTML(200, "login.html", gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("user", username)
	_ = session.Save()

	c.Redirect(http.StatusSeeOther, "/")
}

func (ctrl *AuthController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	_ = session.Save()

	c.Redirect(http.StatusSeeOther, "/")
}
