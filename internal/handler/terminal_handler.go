// internal/handler/terminal_handler.go

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/service"
)

type TerminalHandler struct {
	terminalService service.TerminalService
}

func NewTerminalHandler(terminalService service.TerminalService) *TerminalHandler {
	return &TerminalHandler{
		terminalService: terminalService,
	}
}

func (h *TerminalHandler) Create(w http.ResponseWriter, r *http.Request) {
	var terminal models.TerminalCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&terminal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTerminal, err := h.terminalService.Create(r.Context(), &terminal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTerminal)
}

func (h *TerminalHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid terminal ID", http.StatusBadRequest)
		return
	}

	terminal, err := h.terminalService.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(terminal)
}

func (h *TerminalHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid terminal ID", http.StatusBadRequest)
		return
	}

	var terminal models.TerminalUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&terminal); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTerminal, err := h.terminalService.Update(r.Context(), id, &terminal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTerminal)
}

func (h *TerminalHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid terminal ID", http.StatusBadRequest)
		return
	}

	err = h.terminalService.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *TerminalHandler) List(w http.ResponseWriter, r *http.Request) {
	terminals, err := h.terminalService.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(terminals)
}

func (h *TerminalHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.Create)
	r.Get("/{id}", h.GetByID)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
	r.Get("/", h.List)
	return r
}
