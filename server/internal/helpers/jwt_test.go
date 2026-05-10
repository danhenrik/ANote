package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"anote/internal/constants"
	"anote/internal/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

const testSecret = "test-secret-for-unit-tests"

func setupSecret(t *testing.T) {
	t.Helper()
	constants.JWT_SECRET = testSecret
}

func decodeSegment(t *testing.T, segment string) []byte {
	t.Helper()
	b, err := base64.RawURLEncoding.DecodeString(segment)
	if err != nil {
		t.Fatalf("failed to base64-decode JWT segment: %v", err)
	}
	return b
}

func TestJwtHelper_CreateToken_ProducesValidHS256(t *testing.T) {
	setupSecret(t)
	helper := NewJwtProvider()
	user := &domain.User{Id: "user-123"}

	tokenStr, err := helper.CreateToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenStr)

	parts := strings.Split(tokenStr, ".")
	assert.Equal(t, 3, len(parts), "JWT must have three segments")

	headerBytes := decodeSegment(t, parts[0])
	var header map[string]interface{}
	assert.NoError(t, json.Unmarshal(headerBytes, &header))
	assert.Equal(t, "HS256", header["alg"])
	assert.Equal(t, "JWT", header["typ"])

	payloadBytes := decodeSegment(t, parts[1])
	var payload map[string]interface{}
	assert.NoError(t, json.Unmarshal(payloadBytes, &payload))
	assert.Equal(t, "user-123", payload["id"])

	expRaw, ok := payload["exp"]
	assert.True(t, ok, "exp claim must be present")
	expFloat, ok := expRaw.(float64)
	assert.True(t, ok, "exp claim must be numeric")
	assert.Greater(t, int64(expFloat), time.Now().Unix(), "exp must be in the future")
}

func TestJwtHelper_ValidateToken_AcceptsValidToken(t *testing.T) {
	setupSecret(t)
	helper := NewJwtProvider()
	user := &domain.User{Id: "user-123"}

	tokenStr, err := helper.CreateToken(user)
	assert.NoError(t, err)

	ok, userId := helper.ValidateToken(tokenStr)
	assert.True(t, ok)
	assert.Equal(t, "user-123", userId)
}

func TestJwtHelper_ValidateToken_RejectsAlgNone(t *testing.T) {
	setupSecret(t)
	helper := NewJwtProvider()

	header := map[string]string{"alg": "none", "typ": "JWT"}
	payload := map[string]interface{}{
		"id":  "user-123",
		"exp": time.Now().Add(time.Hour).Unix(),
	}
	headerJSON, _ := json.Marshal(header)
	payloadJSON, _ := json.Marshal(payload)
	forged := base64.RawURLEncoding.EncodeToString(headerJSON) + "." +
		base64.RawURLEncoding.EncodeToString(payloadJSON) + "."

	ok, userId := helper.ValidateToken(forged)
	assert.False(t, ok, "alg=none tokens must be rejected")
	assert.Equal(t, "", userId)
}

func TestJwtHelper_ValidateToken_RejectsAlgConfusion_RS256(t *testing.T) {
	setupSecret(t)
	helper := NewJwtProvider()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(t, err)

	claims := jwt.MapClaims{
		"id":  "user-123",
		"exp": time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signed, err := token.SignedString(privateKey)
	assert.NoError(t, err)

	ok, userId := helper.ValidateToken(signed)
	assert.False(t, ok, "RS256 tokens must be rejected when verifier expects HS256")
	assert.Equal(t, "", userId)
}

func TestJwtHelper_ValidateToken_RejectsExpiredToken(t *testing.T) {
	setupSecret(t)
	helper := NewJwtProvider()

	claims := jwt.MapClaims{
		"id":  "user-123",
		"exp": time.Now().Add(-1 * time.Hour).Unix(),
		"iat": time.Now().Add(-2 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(constants.JWT_SECRET))
	assert.NoError(t, err)

	ok, userId := helper.ValidateToken(signed)
	assert.False(t, ok, "expired tokens must be rejected")
	assert.Equal(t, "", userId)
}

func TestJwtHelper_ValidateToken_RejectsWrongSecret(t *testing.T) {
	setupSecret(t)
	helper := NewJwtProvider()

	claims := jwt.MapClaims{
		"id":  "user-123",
		"exp": time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte("a-different-secret"))
	assert.NoError(t, err)

	ok, userId := helper.ValidateToken(signed)
	assert.False(t, ok, "tokens signed with a different secret must be rejected")
	assert.Equal(t, "", userId)
}

func TestJwtHelper_ValidateToken_RejectsTokenWithoutIdClaim(t *testing.T) {
	setupSecret(t)
	helper := NewJwtProvider()

	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(constants.JWT_SECRET))
	assert.NoError(t, err)

	ok, userId := helper.ValidateToken(signed)
	assert.False(t, ok, "tokens without an id claim must be rejected")
	assert.Equal(t, "", userId)
}

func TestJwtHelper_ValidateToken_RejectsMalformedString(t *testing.T) {
	setupSecret(t)
	helper := NewJwtProvider()

	cases := []string{"", "not.a.jwt", "a.b.c"}
	for _, tc := range cases {
		t.Run(tc, func(t *testing.T) {
			ok, userId := helper.ValidateToken(tc)
			assert.False(t, ok)
			assert.Equal(t, "", userId)
		})
	}
}
