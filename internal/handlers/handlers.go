package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/giovannymassuia/mini-url/internal/dto"
	"github.com/giovannymassuia/mini-url/internal/services"
)

type Handler struct {
	linkService services.LinkServiceInterface
}

func NewHandler(linkService services.LinkServiceInterface) *Handler {
	return &Handler{
		linkService: linkService,
	}
}

func (h *Handler) GetLongLink(w http.ResponseWriter, r *http.Request) {
	shortLink := r.PathValue("shortLink")
	longLink, _ := h.linkService.GetLongUrl(shortLink)
	url := fmt.Sprintf("http://localhost:8080/%s", longLink)

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func (h *Handler) CreateShortLink(w http.ResponseWriter, r *http.Request) {
	var req dto.CrateShortLinkRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	shortUrl, _ := h.linkService.CreateShortLink(req.LongUrl)
	response := dto.CreateShortLinkResponse{ShortUrl: shortUrl}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetAllLinks(w http.ResponseWriter, r *http.Request) {
	links, _ := h.linkService.GetAllLinks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(links)
}
