package main

import (
	"net/http"

	"github.com/giovannymassuia/mini-url/internal/handlers"
	"github.com/giovannymassuia/mini-url/internal/services"
)

type Link struct {
	Id       string `json:"id"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

func main() {
	mux := http.NewServeMux()

	linkService := services.NewLinkService()
	handler := handlers.NewHandler(linkService)

	mux.HandleFunc("GET /api/link", handler.GetAllLinks)
	mux.HandleFunc("GET /{shortLink}", handler.GetLongLink)
	mux.HandleFunc("POST /api/link", handler.CreateShortLink)

	http.ListenAndServe(":8080", mux)
}
