// controllers/register_controller.go

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ozcankasal/potionhub/models"
	"github.com/ozcankasal/potionhub/services"
)

type RegisterController struct {
	userService *services.UserService
}

func NewRegisterController(userService *services.UserService) *RegisterController {
	return &RegisterController{
		userService: userService,
	}
}

func (ctrl *RegisterController) GetRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func (ctrl *RegisterController) PostRegister(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")

	if password != confirmPassword {
		c.HTML(200, "register.html", gin.H{
			"error": "Passwords do not match",
		})
		return
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	err := ctrl.userService.CreateUser(user)
	if err != nil {
		c.HTML(200, "register.html", gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/login")
}
