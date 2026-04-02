package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Ayushmangit/goLangBackendJWT/internal/pkg/jwt"
	"github.com/Ayushmangit/goLangBackendJWT/internal/pkg/response"
	"github.com/Ayushmangit/goLangBackendJWT/internal/pkg/utils"
	"github.com/Ayushmangit/goLangBackendJWT/internal/service"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token,omitempty"`
}

type AuthHandler struct {
	userService service.UserService
}

func NewAuthHandler(s service.UserService) *AuthHandler {
	return &AuthHandler{
		userService: s,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid Request")
		return
	}
	if req.Email == "" || req.Password == "" {
		response.Error(w, http.StatusBadRequest, "email and password required")
		return
	}

	user, err := h.userService.Register(r.Context(), req.Email, req.Password)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusCreated, AuthResponse{
		ID:    user.ID.String(),
		Email: user.Email,
		Role:  user.Role,
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid Request")
		return
	}
	user, err := h.userService.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := jwt.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to generate Token")
	}
	res := AuthResponse{
		ID:    user.ID.String(),
		Email: user.Email,
		Role:  user.Role,
		Token: token,
	}
	response.JSON(w, http.StatusOK, res)
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	authUser, err := utils.GetUserFromContext(r.Context())
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	user, err := h.userService.GetByID(r.Context(), authUser.ID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, AuthResponse{
		ID:    user.ID.String(),
		Email: user.Email,
		Role:  user.Role,
	})
}
