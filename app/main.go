package main

import (
	"net/http"

	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/giovannymassuia/mini-url/docs"
	"github.com/giovannymassuia/mini-url/internal/handlers"
	"github.com/giovannymassuia/mini-url/internal/services"
)

type Link struct {
	Id       string `json:"id"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

// @title Mini URL API
// @version 1.0
// @description This is a simple API to create short links

// @contact.name Giovanny Massuia
// @contact.url giovannymassuia.io

// @host http://localhost:8080
// @BasePath /
// @servers http://localhost:8080
func main() {
	mux := http.NewServeMux()

	linkService := services.NewLinkService()
	handler := handlers.NewHandler(linkService)

	mux.HandleFunc("GET /api/link", handler.GetAllLinks)
	mux.HandleFunc("GET /{shortLink}", handler.GetLongLink)
	mux.HandleFunc("POST /api/link", handler.CreateShortLink)

	// swagger docs
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	// cors
	corsHandler := cors.Default().Handler(mux)

	http.ListenAndServe(":8080", corsHandler)
}
