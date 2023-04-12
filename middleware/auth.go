package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ozcankasal/potionhub/services"
	"golang.org/x/crypto/bcrypt"
)

func AuthRequired(userService *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		givenUsername := c.PostForm("username")
		givenPassword := c.PostForm("password")

		user, err := userService.GetUserByUsername(givenUsername)
		if err == nil && user != nil {
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(givenPassword))
			if err == nil {
				session := sessions.Default(c)
				session.Set("user", givenUsername)
				session.Save()
				c.Next()
			} else {
				c.HTML(401, "login.html", gin.H{
					"title": "Login",
					"error": "Invalid username or password",
				})
				c.Abort()
			}
		} else {
			c.HTML(401, "login.html", gin.H{
				"title": "Login",
				"error": "Invalid username or password",
			})
			c.Abort()
		}
	}
}
