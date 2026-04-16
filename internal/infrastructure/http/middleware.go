package http

import (
	"auth/internal/domain"
	"context"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "userID"

// JWTMiddleware es un decorador de handlers que valida el token antes de seguir.
func JWTMiddleware(tp domain.TokenProvider, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Obtener el header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Capitán, falta el header de autorización", http.StatusUnauthorized)
			return
		}

		// 2. Extraer el token (Formato: Bearer <token>)
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Formato de token inválido. Usá: Bearer <token>", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		// 3. Validar el token usando el Puerto del dominio
		userID, err := tp.Validate(tokenString)
		if err != nil {
			http.Error(w, "Token trucho o vencido, hermano", http.StatusUnauthorized)
			return
		}

		// 4. Inyectar el userID en el contexto para que el handler lo pueda usar
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
