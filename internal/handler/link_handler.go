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

// @Summary Create a new link
// @Description Create a new link between fiscal number and factory number
// @Tags links
// @Accept  json
// @Produce  json
// @Param link body models.LinkCreateRequest true "Create link request"
// @Success 201 {object} models.Link
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security ApiKeyAuth
// @Router /links [post]
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

// @Summary Get a link by factory number
// @Description Get details of a link by its factory number
// @Tags links
// @Accept  json
// @Produce  json
// @Param factoryNumber path string true "Factory Number"
// @Success 200 {object} models.Link
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security ApiKeyAuth
// @Router /links/factory/{factoryNumber} [get]
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

// @Summary Delete a link
// @Description Delete a link by its ID
// @Tags links
// @Accept  json
// @Produce  json
// @Param id path int true "Link ID"
// @Success 204 "No Content"
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Security ApiKeyAuth
// @Router /links/{id} [delete]
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

// @Summary List all links
// @Description Get a list of all links
// @Tags links
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Link
// @Failure 500 {object} models.ErrorResponse
// @Security ApiKeyAuth
// @Router /links [get]
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
