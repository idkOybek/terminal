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

type FiscalModuleHandler struct {
	service *service.FiscalModuleService
	logger  *logger.Logger
}

func NewFiscalModuleHandler(service *service.FiscalModuleService, logger *logger.Logger) *FiscalModuleHandler {
	return &FiscalModuleHandler{
		service: service,
		logger:  logger,
	}
}

// @Security Bearer
// @Summary Create a new fiscal module
// @Description Create a new fiscal module with the given input
// @Tags fiscal-modules
// @Accept  json
// @Produce  json
// @Param fiscal_module body models.FiscalModuleCreateRequest true "Create fiscal module request"
// @Success 201 {object} models.FiscalModule
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /fiscal-modules [post]
func (h *FiscalModuleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req models.FiscalModuleCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Failed to decode request body", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	module, err := h.service.Create(r.Context(), &req)
	if err != nil {
		h.logger.Error("Failed to create fiscal module", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to create fiscal module")
		return
	}

	RespondWithJSON(w, http.StatusCreated, module)
}

// @Security Bearer
// @Summary Get a fiscal module by ID
// @Description Get details of a fiscal module by its ID
// @Tags fiscal-modules
// @Accept  json
// @Produce  json
// @Param id path int true "Fiscal Module ID"
// @Success 200 {object} models.FiscalModule
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /fiscal-modules/{id} [get]
func (h *FiscalModuleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("Invalid fiscal module ID", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid fiscal module ID")
		return
	}

	module, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		h.logger.Error("Failed to get fiscal module", "error", err)
		RespondWithError(w, http.StatusNotFound, "Fiscal module not found")
		return
	}

	RespondWithJSON(w, http.StatusOK, module)
}

// @Security Bearer
// @Summary Update a fiscal module
// @Description Update a fiscal module's details by its ID
// @Tags fiscal-modules
// @Accept  json
// @Produce  json
// @Param id path int true "Fiscal Module ID"
// @Param fiscal_module body models.FiscalModuleUpdateRequest true "Update fiscal module request"
// @Success 200 {object} models.FiscalModule
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /fiscal-modules/{id} [put]
func (h *FiscalModuleHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("Invalid fiscal module ID", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid fiscal module ID")
		return
	}

	var req models.FiscalModuleUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Failed to decode request body", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	module, err := h.service.Update(r.Context(), id, &req)
	if err != nil {
		h.logger.Error("Failed to update fiscal module", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to update fiscal module")
		return
	}

	RespondWithJSON(w, http.StatusOK, module)
}

// @Security Bearer
// @Summary Delete a fiscal module
// @Description Delete a fiscal module by its ID
// @Tags fiscal-modules
// @Accept  json
// @Produce  json
// @Param id path int true "Fiscal Module ID"
// @Success 204 "No Content"
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /fiscal-modules/{id} [delete]
func (h *FiscalModuleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("Invalid fiscal module ID", "error", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid fiscal module ID")
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		h.logger.Error("Failed to delete fiscal module", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to delete fiscal module")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Security Bearer
// @Summary List all fiscal modules
// @Description Get a list of all fiscal modules
// @Tags fiscal-modules
// @Accept  json
// @Produce  json
// @Success 200 {array} models.FiscalModule
// @Failure 500 {object} models.ErrorResponse
// @Router /fiscal-modules [get]
func (h *FiscalModuleHandler) List(w http.ResponseWriter, r *http.Request) {
	modules, err := h.service.List(r.Context())
	if err != nil {
		h.logger.Error("Failed to fetch fiscal modules", "error", err)
		RespondWithError(w, http.StatusInternalServerError, "Failed to fetch fiscal modules")
		return
	}

	RespondWithJSON(w, http.StatusOK, modules)
}

func (h *FiscalModuleHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.Create)
	r.Get("/{id}", h.GetByID)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
	r.Get("/", h.List)
	return r
}
