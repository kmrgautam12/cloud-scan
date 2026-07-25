package auth

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
)

const authHeaderPrefix string = "Bearer"

type AuthMiddlewareService struct {
	middlewareService JWTService
}

func NewMiddlewareService(ts JWTService) *AuthMiddlewareService {
	return &AuthMiddlewareService{
		middlewareService: ts,
	}
}

func (a *AuthMiddlewareService) ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := extractBearerToken(c)
		if err != nil {
			return echo.NewHTTPError(401)
		}
	}

}

func extractBearerToken(c echo.Context) (string, error) {

	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}
	if !strings.HasPrefix(authHeader, authHeaderPrefix) {
		return "", errors.New("Invalid authorization header")
	}
	token := strings.TrimPrefix(authHeader, authHeaderPrefix)
	if token == "" {
		return "", errors.New("Invalid authorization token")
	}
	return token, nil
}
