package middlewares

import (
	"anote/internal/constants"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func newRouterWithSecurityHeaders() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(SecurityHeaders())
	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })
	return r
}

func performGet(r *gin.Engine, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	r.ServeHTTP(w, req)
	return w
}

func TestSecurityHeaders_AppliesBaseHeaders(t *testing.T) {
	previous := constants.SECURE_COOKIE
	t.Cleanup(func() { constants.SECURE_COOKIE = previous })
	constants.SECURE_COOKIE = false

	r := newRouterWithSecurityHeaders()
	w := performGet(r, "/ping")

	csp := w.Header().Get("Content-Security-Policy")
	assert.NotEmpty(t, csp, "CSP header should be set")
	assert.True(t, strings.Contains(csp, "default-src 'self'"), "CSP should contain default-src 'self'")

	assert.Equal(t, "nosniff", w.Header().Get("X-Content-Type-Options"))
	assert.Equal(t, "no-referrer", w.Header().Get("Referrer-Policy"))
	assert.Equal(t, "DENY", w.Header().Get("X-Frame-Options"))
}

func TestSecurityHeaders_HSTSOnlyWhenSecureCookie(t *testing.T) {
	previous := constants.SECURE_COOKIE
	t.Cleanup(func() { constants.SECURE_COOKIE = previous })

	constants.SECURE_COOKIE = false
	r := newRouterWithSecurityHeaders()
	w := performGet(r, "/ping")
	assert.Empty(t, w.Header().Get("Strict-Transport-Security"), "HSTS should be absent when SECURE_COOKIE=false")

	constants.SECURE_COOKIE = true
	r = newRouterWithSecurityHeaders()
	w = performGet(r, "/ping")
	assert.NotEmpty(t, w.Header().Get("Strict-Transport-Security"), "HSTS should be present when SECURE_COOKIE=true")
}

func TestSecurityHeaders_AppliedOnNotFound(t *testing.T) {
	previous := constants.SECURE_COOKIE
	t.Cleanup(func() { constants.SECURE_COOKIE = previous })
	constants.SECURE_COOKIE = false

	r := newRouterWithSecurityHeaders()
	w := performGet(r, "/does-not-exist")

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.NotEmpty(t, w.Header().Get("Content-Security-Policy"), "security headers must apply to 404 responses too")
}
