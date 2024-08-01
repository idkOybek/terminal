package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/service"
	"github.com/idkOybek/newNewTerminal/pkg/logger"
)

type AuthHandler struct {
	service *service.AuthService
	logger  *logger.Logger
}

func NewAuthHandler(service *service.AuthService, logger *logger.Logger) *AuthHandler {
	return &AuthHandler{
		service: service,
		logger:  logger,
	}
}

// @Security Bearer
// @Summary Register a new user
// @Description Register a new user with the given input
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.UserCreateRequest true "User registration info"
// @Success 201 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Failed to decode request body", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := h.service.Register(r.Context(), &req)
	if err != nil {
		h.logger.Error("Failed to register user", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to register user")
		return
	}

	RespondWithJSON(w, http.StatusCreated, user)
}

// @Summary Login user
// @Description Authenticate a user and return a token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param credentials body models.UserLoginRequest true "User login credentials"
// @Success 200 {object} models.UserLoginResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Failed to decode request body", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	resp, err := h.service.Login(r.Context(), &req)
	if err != nil {
		h.logger.Error("Failed to login user", "error", err)
		RespondWithError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	RespondWithJSON(w, http.StatusOK, resp)
}

func (h *AuthHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/register", h.Register)
	r.Post("/login", h.Login)
	return r
}
