package middlewares

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/constants"
	"anote/internal/container"
	"anote/internal/domain"
	"anote/internal/errors"
	"anote/internal/helpers"
	mock_interfaces "anote/internal/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func setupJwtProvider(t *testing.T) helpers.JwtHelper {
	t.Helper()
	constants.JWT_SECRET = "test-secret"
	provider := helpers.NewJwtProvider()
	container.JwtProvider = provider
	return provider
}

func makeRequestWithCookie(token string) httpAdapter.Request {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	if token != "" {
		r.AddCookie(&http.Cookie{Name: "auth_token", Value: token})
	}
	return httpAdapter.Request{
		Method:  r.Method,
		Headers: r.Header,
		Cookies: r.Cookies(),
	}
}

func makeRequestWithHeader(headerValue string) httpAdapter.Request {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("Authorization", headerValue)
	return httpAdapter.Request{
		Method:  r.Method,
		Headers: r.Header,
		Cookies: r.Cookies(),
	}
}

func makeRequestWithCookieAndHeader(token, headerValue string) httpAdapter.Request {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	if token != "" {
		r.AddCookie(&http.Cookie{Name: "auth_token", Value: token})
	}
	if headerValue != "" {
		r.Header.Set("Authorization", headerValue)
	}
	return httpAdapter.Request{
		Method:  r.Method,
		Headers: r.Header,
		Cookies: r.Cookies(),
	}
}

func TestAuthMiddleware_AcceptsCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := setupJwtProvider(t)
	user := &domain.User{Id: "alice", Email: "alice@example.com"}
	token, err := provider.CreateToken(user)
	assert.NoError(t, err)

	userRepo := mock_interfaces.NewMockUserRepository(ctrl)
	userRepo.EXPECT().GetByUsername("alice").Return(user, nil)
	container.UserRepository = userRepo

	req := makeRequestWithCookie(token)
	out, appErr := AuthenticateUser(req)

	assert.Nil(t, appErr)
	assert.Equal(t, "alice", out.User.ID)
	assert.Equal(t, "alice@example.com", out.User.Email)
}

func TestAuthMiddleware_AcceptsBearerHeader_BackwardCompat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := setupJwtProvider(t)
	user := &domain.User{Id: "bob", Email: "bob@example.com"}
	token, err := provider.CreateToken(user)
	assert.NoError(t, err)

	userRepo := mock_interfaces.NewMockUserRepository(ctrl)
	userRepo.EXPECT().GetByUsername("bob").Return(user, nil)
	container.UserRepository = userRepo

	req := makeRequestWithHeader("Bearer " + token)
	out, appErr := AuthenticateUser(req)

	assert.Nil(t, appErr)
	assert.Equal(t, "bob", out.User.ID)
}

func TestAuthMiddleware_PrefersCookieOverHeader(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	provider := setupJwtProvider(t)
	cookieUser := &domain.User{Id: "from-cookie", Email: "cookie@example.com"}
	headerUser := &domain.User{Id: "from-header", Email: "header@example.com"}

	cookieToken, err := provider.CreateToken(cookieUser)
	assert.NoError(t, err)
	headerToken, err := provider.CreateToken(headerUser)
	assert.NoError(t, err)

	userRepo := mock_interfaces.NewMockUserRepository(ctrl)
	userRepo.EXPECT().GetByUsername("from-cookie").Return(cookieUser, nil)
	container.UserRepository = userRepo

	req := makeRequestWithCookieAndHeader(cookieToken, "Bearer "+headerToken)
	out, appErr := AuthenticateUser(req)

	assert.Nil(t, appErr)
	assert.Equal(t, "from-cookie", out.User.ID)
}

func TestAuthMiddleware_Returns401WhenNeitherPresent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	setupJwtProvider(t)
	userRepo := mock_interfaces.NewMockUserRepository(ctrl)
	container.UserRepository = userRepo

	req := makeRequestWithCookie("")
	_, appErr := AuthenticateUser(req)

	assert.NotNil(t, appErr)
	assert.Equal(t, uint(401), appErr.Status)
}

func TestAuthMiddleware_Returns400OnInvalidCookie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	setupJwtProvider(t)
	userRepo := mock_interfaces.NewMockUserRepository(ctrl)
	container.UserRepository = userRepo

	req := makeRequestWithCookie("not-a-real-jwt")
	_, appErr := AuthenticateUser(req)

	assert.NotNil(t, appErr)
	assert.Equal(t, uint(400), appErr.Status)
	assert.Equal(t, "Invalid JWT Token", appErr.Message)
}

func TestAuthMiddleware_DoesNotPanicOnMalformedAuthHeader(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	setupJwtProvider(t)
	userRepo := mock_interfaces.NewMockUserRepository(ctrl)
	container.UserRepository = userRepo

	cases := []string{"Bearer", "Bearer ", "Token abc"}
	for _, hv := range cases {
		req := makeRequestWithHeader(hv)
		var appErr *errors.AppError
		assert.NotPanics(t, func() {
			_, appErr = AuthenticateUser(req)
		})
		assert.NotNil(t, appErr, "header value %q should yield error", hv)
	}
}
