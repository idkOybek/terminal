// internal/handler/fiscal_module_handler.go

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/service"
)

type FiscalModuleHandler struct {
	fiscalModuleService service.FiscalModuleService
}

func NewFiscalModuleHandler(fiscalModuleService service.FiscalModuleService) *FiscalModuleHandler {
	return &FiscalModuleHandler{
		fiscalModuleService: fiscalModuleService,
	}
}

// @Summary Create a new fiscal module
// @Description Create a new fiscal module with the given input
// @Tags fiscal-modules
// @Accept  json
// @Produce  json
// @Param module body models.FiscalModuleCreateRequest true "Create fiscal module request"
// @Success 201 {object} models.FiscalModule
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security ApiKeyAuth
// @Router /fiscal-modules [post]
func (h *FiscalModuleHandler) Create(w http.ResponseWriter, r *http.Request) {
	var module models.FiscalModuleCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&module); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdModule, err := h.fiscalModuleService.Create(r.Context(), &module)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdModule)
}

// @Summary Get a fiscal module by ID
// @Description Get details of a fiscal module by its ID
// @Tags fiscal-modules
// @Accept  json
// @Produce  json
// @Param id path int true "Fiscal Module ID"
// @Success 200 {object} models.FiscalModule
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security ApiKeyAuth
// @Router /fiscal-modules/{id} [get]
func (h *FiscalModuleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid fiscal module ID", http.StatusBadRequest)
		return
	}

	module, err := h.fiscalModuleService.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(module)
}

// @Summary Update a fiscal module
// @Description Update a fiscal module's details by its ID
// @Tags fiscal-modules
// @Accept  json
// @Produce  json
// @Param id path int true "Fiscal Module ID"
// @Param module body models.FiscalModuleUpdateRequest true "Update fiscal module request"
// @Success 200 {object} models.FiscalModule
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security ApiKeyAuth
// @Router /fiscal-modules/{id} [put]
func (h *FiscalModuleHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid fiscal module ID", http.StatusBadRequest)
		return
	}

	var module models.FiscalModuleUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&module); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedModule, err := h.fiscalModuleService.Update(r.Context(), id, &module)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedModule)
}

// @Summary Delete a fiscal module
// @Description Delete a fiscal module by its ID
// @Tags fiscal-modules
// @Accept  json
// @Produce  json
// @Param id path int true "Fiscal Module ID"
// @Success 204 "No Content"
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security ApiKeyAuth
// @Router /fiscal-modules/{id} [delete]
func (h *FiscalModuleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid fiscal module ID", http.StatusBadRequest)
		return
	}

	err = h.fiscalModuleService.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary List all fiscal modules
// @Description Get a list of all fiscal modules
// @Tags fiscal-modules
// @Accept  json
// @Produce  json
// @Success 200 {array} models.FiscalModule
// @Failure 500 {object} models.ErrorResponse
// @Security ApiKeyAuth
// @Router /fiscal-modules [get]
func (h *FiscalModuleHandler) List(w http.ResponseWriter, r *http.Request) {
	modules, err := h.fiscalModuleService.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modules)
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
