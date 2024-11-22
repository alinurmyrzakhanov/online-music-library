package server

import (
	"database/sql"
	"log"
	"net/http"

	_ "online-music-library/docs"
	"online-music-library/internal/clients"
	handlers "online-music-library/internal/command"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewServer(db *sql.DB, musicClient *clients.MusicInfoClient) http.Handler {
	r := chi.NewRouter()
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}
	if musicClient == nil {
		log.Fatal("MusicInfoClient is not initialized")
	}
	h := handlers.NewHandlers(db, musicClient)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))

	// Routes
	r.Get("/songs", h.GetSongs)           // @Summary Get all songs
	r.Get("/songs/{id}", h.GetSongByID)   // @Summary Get song by ID
	r.Post("/songs", h.AddSong)           // @Summary Add a new song
	r.Put("/songs/{id}", h.UpdateSong)    // @Summary Update a song
	r.Delete("/songs/{id}", h.DeleteSong) // @Summary Delete a song

	// Swagger route
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
