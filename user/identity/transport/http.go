package transport

import (
	"net/http"

	"github.com/RagOfJoes/idp/session"
	"github.com/RagOfJoes/idp/transport"
	"github.com/gin-gonic/gin"
)

var (
	errNotAuthenticated error = transport.NewHttpClientError(http.StatusUnauthorized, "not_authenticated", "Not authenticated", nil)
)

type Http struct {
	sm *session.Manager
}

func NewIdentityHttp(sm *session.Manager, r *gin.Engine) {
	h := &Http{
		sm: sm,
	}
	r.GET("/me", h.me())
}

func (h *Http) me() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := transport.IsAuthenticated(c)
		if sess == nil {
			c.Error(errNotAuthenticated)
			return
		}

		c.JSON(http.StatusOK, transport.HttpResponse{
			Success: true,
			Payload: sess,
		})
	}
}
