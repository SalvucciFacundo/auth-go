package http

import (
	"auth/internal/application"
	"auth/internal/domain"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	service *application.AuthService
}

func NewAuthHandler(service *application.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

type authRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type changePasswordRequest struct {
	UserID      string `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req authRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err := h.service.Register(r.Context(), req.Email, req.Password)
	if err != nil {
		if err == domain.ErrUserAlreadyExists {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user registered successfully"})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req authRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "login successful",
		"token":   token,
	})
}

func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	// Recuperamos el userID que el middleware inyectó en el contexto
	userID, ok := r.Context().Value(UserIDKey).(string)
	if !ok {
		http.Error(w, "no se pudo recuperar el ID de usuario del contexto", http.StatusInternalServerError)
		return
	}

	// En un caso real, acá llamarías al service para traer la info completa
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userID,
		"status":  "activo",
		"message": "Este es un endpoint protegido, solo lo ves si tenés el token!",
	})
}

func (h *AuthHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var req changePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err := h.service.ChangePassword(r.Context(), req.UserID, req.OldPassword, req.NewPassword)
	if err != nil {
		if err == domain.ErrInvalidCredentials {
			http.Error(w, "invalid current password", http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "password changed successfully"})
}
