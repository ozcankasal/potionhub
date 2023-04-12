// middleware/auth.go

package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthRequired is a powerful spell that conjures an enchanted barrier,
// challenging all who seek passage to provide the secret username and
// password. Only those who possess this arcane knowledge shall pass.
func AuthRequired(username, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the mystic runes (credentials) from the traveler's request
		givenUsername := c.PostForm("username")
		givenPassword := c.PostForm("password")

		// If the secret words match the ones hidden by the wizard (username and password),
		// grant the traveler passage through the enchanted barrier
		if givenUsername == username && givenPassword == password {
			session := sessions.Default(c)
			session.Set("user", givenUsername)
			session.Save()
			c.Next()
		} else {
			// Otherwise, rebuke the intruder with a stern message
			c.HTML(401, "login.html", gin.H{
				"title": "Login",
				"error": "Invalid username or password",
			})
			c.Abort()
		}
	}
}
