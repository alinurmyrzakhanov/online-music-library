package main

import (
	"log"
	"net/http"

	"online-music-library/internal/clients"
	"online-music-library/internal/config"
	"online-music-library/internal/repositories"
	"online-music-library/internal/server"
)

func main() {
	cfg := config.LoadConfig()

	db, err := repositories.NewDB(cfg)
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	musicClient := clients.NewMusicInfoClient(cfg.MusicAPI)

	srv := server.NewServer(db, musicClient)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
