package middlewares

import (
	"anote/internal/constants"

	"github.com/gin-gonic/gin"
)

// SecurityHeaders returns a Gin middleware that applies a baseline set of
// security headers (CSP, X-Content-Type-Options, Referrer-Policy,
// X-Frame-Options) to every response. HSTS is added only when the runtime
// is configured to expect HTTPS via SECURE_COOKIE.
func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header(
			"Content-Security-Policy",
			"default-src 'self'; img-src 'self' data:; style-src 'self' 'unsafe-inline'; script-src 'self'",
		)
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Referrer-Policy", "no-referrer")
		c.Header("X-Frame-Options", "DENY")
		if constants.SECURE_COOKIE {
			c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		}
		c.Next()
	}
}
