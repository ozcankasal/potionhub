package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired(username, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		givenUsername := c.PostForm("username")
		givenPassword := c.PostForm("password")

		if givenUsername == username && givenPassword == password {
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
	}
}
