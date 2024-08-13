package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/service"
	"github.com/idkOybek/newNewTerminal/pkg/logger"
)

type UserHandler struct {
	service *service.UserService
	logger  *logger.Logger
}

func NewUserHandler(service *service.UserService, logger *logger.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

// @Security Bearer
// @Summary Create a new user
// @Description Create a new user with the given input
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.UserCreateRequest true "Create user request"
// @Success 201 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Failed to decode request body", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := h.service.Create(r.Context(), &req)
	if err != nil {
		h.logger.Error("Failed to create user", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	RespondWithJSON(w, http.StatusCreated, user)
}

// @Security Bearer
// @Summary Get a user by ID
// @Description Get details of a user by its ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("Invalid user ID", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		h.logger.Error("Failed to get user", "error", err)
		RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	RespondWithJSON(w, http.StatusOK, user)
}

// @Security Bearer
// @Summary Update a user
// @Description Update a user's details by its ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body models.UserUpdateRequest true "Update user request"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [put]
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("Invalid user ID", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req models.UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Failed to decode request body", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := h.service.Update(r.Context(), id, &req)
	if err != nil {
		h.logger.Error("Failed to update user", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	RespondWithJSON(w, http.StatusOK, user)
}

// @Security Bearer
// @Summary Delete a user
// @Description Delete a user by its ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{id} [delete]
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("Invalid user ID", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		h.logger.Error("Failed to delete user and associated data", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to delete user and associated data")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Security Bearer
// @Summary List all users
// @Description Get a list of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Failure 500 {object} models.ErrorResponse
// @Router /users [get]
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.List(r.Context())
	if err != nil {
		h.logger.Error("Failed to fetch users", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	RespondWithJSON(w, http.StatusOK, users)
}

func (h *UserHandler) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Post("/", h.Create)
	r.Get("/{id}", h.GetByID)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
	r.Get("/", h.List)
	return r
}
