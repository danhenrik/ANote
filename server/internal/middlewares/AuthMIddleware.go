package middlewares

import (
	httpAdapter "anote/cmd/interfaces"
	"anote/internal/container"
	"anote/internal/errors"
	"log"
	"strings"
)

const authCookieName = "auth_token"

func AuthenticateUser(req httpAdapter.Request) (httpAdapter.Request, *errors.AppError) {
	jwtToken := extractTokenFromCookie(req)
	if jwtToken == "" {
		headerToken, ok := extractTokenFromHeader(req)
		if !ok {
			log.Println("[AuthMiddleware] No auth credentials found in cookie or header")
			return req, errors.NewAppError(401, "Couldn't locate a valid Authorization Header")
		}
		jwtToken = headerToken
	}

	tokenIsValid, userId := container.JwtProvider.ValidateToken(jwtToken)
	if !tokenIsValid {
		log.Println("[AuthMiddleware] Invalid JWT Token")
		return req, errors.NewAppError(400, "Invalid JWT Token")
	}

	user, err := container.UserRepository.GetByUsername(userId)
	if err != nil {
		log.Println("[AuthMiddleware] Error on get token owner", err)
		return req, err
	}

	req.User = httpAdapter.UserIdentity{
		ID:    user.Id,
		Email: user.Email,
	}
	log.Println("[AuthMiddleware] Successfully authenticated user", user.Id)
	return req, nil
}

func extractTokenFromCookie(req httpAdapter.Request) string {
	for _, c := range req.Cookies {
		if c.Name == authCookieName {
			return c.Value
		}
	}
	return ""
}

func extractTokenFromHeader(req httpAdapter.Request) (string, bool) {
	bearerToken, found := req.GetHeader("Authorization")
	if !found {
		return "", false
	}
	parts := strings.SplitN(bearerToken, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") || strings.TrimSpace(parts[1]) == "" {
		return "", false
	}
	return strings.TrimSpace(parts[1]), true
}
