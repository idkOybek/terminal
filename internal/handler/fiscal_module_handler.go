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