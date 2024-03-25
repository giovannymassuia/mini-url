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

// @Summary Get long link
// @Description Get long link from short link
// @Tags links
// @Param shortLink path string true "Short link"
// @Produce html
// @Success 301 {string} string "Redirect to long link"
// @Failure 404 {string} string "Link not found"
// @Router /{shortLink} [get]
func (h *Handler) GetLongLink(w http.ResponseWriter, r *http.Request) {
	shortLink := r.PathValue("shortLink")
	longLink, _ := h.linkService.GetLongUrl(shortLink)

	fmt.Println("Redirecting to", longLink)

	// no-cache
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	http.Redirect(w, r, longLink, http.StatusMovedPermanently)
}

// @Summary Create short link
// @Description Create short link from long link
// @Tags links
// @Accept json
// @Produce json
// @Param request body dto.CrateShortLinkRequest true "Long URL"
// @Success 201 {object} dto.CreateShortLinkResponse
// @Header 201 {string} Location "/{shortUrl}"
// @Router /api/link [post]
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
	w.Header().Set("Location", fmt.Sprintf("%s", shortUrl))
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

// @Summary Get all links
// @Description Get all links
// @Tags links
// @Produce json
// @Success 200 {array} domain.Link
// @Router /api/link [get]
func (h *Handler) GetAllLinks(w http.ResponseWriter, r *http.Request) {
	links, _ := h.linkService.GetAllLinks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(links)
}
