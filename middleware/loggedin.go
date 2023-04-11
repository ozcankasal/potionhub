// middleware/loggedin.go

package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		if user != nil {
			c.Next()
		} else {
			c.Redirect(302, "/welcome")
			c.Abort()
		}
	}
}
