// middleware/loggedin.go

package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoggedIn casts a protective enchantment around our mystical routes,
// allowing only those who possess the sacred token to access them.
func LoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Invoke the ancient art of session magic to reveal the user's identity
		session := sessions.Default(c)
		user := session.Get("user")

		// If the sacred token is present, grant passage through the enchanted gate
		if user != nil {
			c.Next()
		} else {
			// Otherwise, banish the trespassers to the Welcome realm
			c.Redirect(302, "/welcome")
			c.Abort()
		}
	}
}
