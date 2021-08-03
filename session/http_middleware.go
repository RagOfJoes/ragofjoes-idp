package session

import (
	"github.com/gin-gonic/gin"
)

// AuthMiddleware retrieves identity from session
// and passes it to context
func AuthMiddleware(sm *Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := sm.GetAuth(c.Request.Context(), true)
		c.Set("auth_session", sess)

		c.Next()
	}
}
