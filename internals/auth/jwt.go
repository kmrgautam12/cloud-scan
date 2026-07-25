package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret []byte
	issuer string
}

func NewJwtService(secret, issuer string) *JWTService {
	return &JWTService{
		secret: []byte(secret),
		issuer: issuer,
	}
}

func (j *JWTService) GenerateToken(username, role string) (string, error) {
	claims := &Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			Subject:   username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString(j.secret)
	if err != nil {
		return "", err
	}
	return tokenstr, nil
}

func (j *JWTService) ValidateToken(token string) (claim Claims, err error) {
	jwt.ParseWithClaims(token, &Claims{})

}
