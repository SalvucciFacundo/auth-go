package main

import (
	"auth/internal/application"
	web "auth/internal/infrastructure/http"
	"auth/internal/infrastructure/postgres"
	"auth/internal/infrastructure/security"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// 1. Configuración de la DB (En un entorno real usarías variables de entorno!)
	dbURL := "postgres://kuno:secret_password@localhost:5432/auth_db"
	
	// Esperar un poquito a que Docker levante si es necesario
	ctx := context.Background()
	
	dbPool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbPool.Close()

	// 2. Inicializar capas (Inyección de dependencias)
	userRepo := postgres.NewUserRepository(dbPool)
	
	// Implementación del puerto TokenProvider
	jwtProvider := security.NewJWTProvider("mi_clave_secreta_super_segura_123")
	
	authService := application.NewAuthService(userRepo, jwtProvider)
	authHandler := web.NewAuthHandler(authService)

	// 3. Definir Rutas
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("POST /login", authHandler.Login)
	mux.HandleFunc("POST /change-password", authHandler.ChangePassword)

	// 3.1 Ruta Protegida: Solo accesible con un token válido.
	// El middleware valida el token usando el provider e inyecta el ID en el request.
	mux.HandleFunc("GET /me", web.JWTMiddleware(jwtProvider, authHandler.GetProfile))

	// 4. Arrancar el servidor
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
