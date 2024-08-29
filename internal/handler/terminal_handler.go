package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/service"
	"github.com/idkOybek/newNewTerminal/pkg/logger"
)

type TerminalHandler struct {
	service *service.TerminalService
	logger  *logger.Logger
}

func NewTerminalHandler(service *service.TerminalService, logger *logger.Logger) *TerminalHandler {
	if service == nil {
		log.Println("Error: TerminalService is nil in NewTerminalHandler")
		return nil
	}
	if logger == nil {
		log.Println("Warning: logger is nil in NewTerminalHandler")
	}

	return &TerminalHandler{
		service: service,
		logger:  logger,
	}
}

// @Security Bearer
// @Summary Create a new terminal
// @Description Create a new terminal with the given input
// @Tags terminals
// @Accept  json
// @Produce  json
// @Success 201 {object} models.Terminal
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /terminals [post]
func (h *TerminalHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.TerminalCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Failed to decode request body", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	terminal, err := h.service.Create(r.Context(), &req)
	if err != nil {
		h.logger.Error("Failed to create terminal", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to create terminal")
		return
	}

	h.logger.Info("Terminal created successfully", "id", terminal.ID)

	RespondWithJSON(w, http.StatusCreated, terminal)
}

// @Security Bearer
// @Summary Get an exists of terminal by CashRegister
// @Description Get an exists of terminal by CashRegister
// @Tags terminals
// @Accept  json
// @Produce  json
// @Param terminal body models.TerminalExistsRequest true "Get an exists request"
// @Success 200 {object} models.TerminalExistsResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /terminals/exists [get]
func (h *TerminalHandler) CheckExists(w http.ResponseWriter, r *http.Request) {
	var req models.TerminalExistsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Failed to decode request body", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	response, err := h.service.CheckExists(r.Context(), req.CashRegisterNumber)
	if err != nil {
		h.logger.Error("Failed to check terminal existence", "error", err)
		RespondWithError(w, http.StatusNotFound, "Terminal not found")
		return
	}

	RespondWithJSON(w, http.StatusOK, response)
}

// @Security Bearer
// @Summary Get a status of terminal by ID
// @Description Get status of terminal by its ID
// @Tags terminals
// @Accept  json
// @Produce  json
// @Param id path int true "Terminal ID"
// @Success 200 {object} models.TerminalStatusResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /terminals/status/{id} [get]
func (h *TerminalHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error("Invalid terminal ID", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid terminal ID")
		return
	}

	status, err := h.service.GetStatus(r.Context(), id)
	if err != nil {
		h.logger.Error("Failed to get terminal status", "error", err)
		RespondWithError(w, http.StatusNotFound, "Terminal not found")
		return
	}

	RespondWithJSON(w, http.StatusOK, status)
}

// @Security Bearer
// @Summary Get a terminal by ID
// @Description Get details of a terminal by its ID
// @Tags terminals
// @Accept  json
// @Produce  json
// @Param id path int true "Terminal ID"
// @Success 200 {object} models.Terminal
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /terminals/{id} [get]
func (h *TerminalHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("Invalid terminal ID", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid terminal ID")
		return
	}

	terminal, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		h.logger.Error("Failed to get terminal", "error", err)
		RespondWithError(w, http.StatusNotFound, "Terminal not found")
		return
	}

	RespondWithJSON(w, http.StatusOK, terminal)
}

// @Security Bearer
// @Summary Update a terminal
// @Description Update a terminal's details by its ID
// @Tags terminals
// @Accept  json
// @Produce  json
// @Param id path int true "Terminal ID"
// @Param terminal body models.TerminalUpdateRequest true "Update terminal request"
// @Success 200 {object} models.Terminal
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /terminals/{id} [put]
func (h *TerminalHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("Invalid terminal ID", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid terminal ID")
		return
	}

	var req models.TerminalUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Failed to decode request body", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	terminal, err := h.service.Update(r.Context(), id, &req)
	if err != nil {
		h.logger.Error("Failed to update terminal", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to update terminal")
		return
	}

	RespondWithJSON(w, http.StatusOK, terminal)
}

// @Security Bearer
// @Summary Delete a terminal
// @Description Delete a terminal by its ID
// @Tags terminals
// @Accept  json
// @Produce  json
// @Param id path int true "Terminal ID"
// @Success 204 "No Content"
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /terminals/{id} [delete]
func (h *TerminalHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("Invalid terminal ID", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid terminal ID")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		h.logger.Error("Failed to delete terminal", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to delete terminal")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Security Bearer
// @Summary List all terminals
// @Description Get a list of all terminals
// @Tags terminals
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Terminal
// @Failure 500 {object} models.ErrorResponse
// @Router /terminals [get]
func (h *TerminalHandler) List(w http.ResponseWriter, r *http.Request) {
	terminals, err := h.service.List(r.Context())
	if err != nil {
		h.logger.Error("Failed to fetch terminals", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to fetch terminals")
		return
	}

	RespondWithJSON(w, http.StatusOK, terminals)
}

func (h *TerminalHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.Create)
	r.Get("/{id}", h.GetByID)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
	r.Get("/", h.List)
	r.Post("/exists", h.CheckExists)
	r.Get("/status/{id}", h.GetStatus)
	return r
}
