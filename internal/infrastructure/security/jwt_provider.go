package security

import (
	"auth/internal/domain"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTProvider es la implementación concreta de TokenProvider usando JWT.
type JWTProvider struct {
	secretKey []byte
}

func NewJWTProvider(secret string) *JWTProvider {
	return &JWTProvider{
		secretKey: []byte(secret),
	}
}

// Generate crea un token firmado para un usuario específico.
func (p *JWTProvider) Generate(user *domain.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // El token dura 24 horas
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(p.secretKey)
}

// Validate verifica que el token sea válido y devuelve el userID (el "sub").
func (p *JWTProvider) Validate(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Validamos que el método de firma sea el que esperamos (HMAC)
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inesperado: %v", t.Header["alg"])
		}
		return p.secretKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("error al parsear token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["sub"].(string)
		if !ok {
			return "", fmt.Errorf("el claim 'sub' no es válido")
		}
		return userID, nil
	}

	return "", fmt.Errorf("token inválido o expirado")
}
