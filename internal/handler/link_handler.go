// internal/handler/link_handler.go

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/idkOybek/newNewTerminal/internal/models"
	"github.com/idkOybek/newNewTerminal/internal/service"
)

type LinkHandler struct {
	linkService service.LinkService
}

func NewLinkHandler(linkService service.LinkService) *LinkHandler {
	return &LinkHandler{
		linkService: linkService,
	}
}

func (h *LinkHandler) Create(w http.ResponseWriter, r *http.Request) {
	var link models.LinkCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdLink, err := h.linkService.Create(r.Context(), &link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdLink)
}

func (h *LinkHandler) GetByFactoryNumber(w http.ResponseWriter, r *http.Request) {
	factoryNumber := chi.URLParam(r, "factoryNumber")

	link, err := h.linkService.GetByFactoryNumber(r.Context(), factoryNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(link)
}

func (h *LinkHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid link ID", http.StatusBadRequest)
		return
	}

	err = h.linkService.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *LinkHandler) List(w http.ResponseWriter, r *http.Request) {
	links, err := h.linkService.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(links)
}

func (h *LinkHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", h.Create)
	r.Get("/factory/{factoryNumber}", h.GetByFactoryNumber)
	r.Delete("/{id}", h.Delete)
	r.Get("/", h.List)
	return r
}