package command

import (
	"database/sql"
	"online-music-library/internal/clients"
)

type Handlers struct {
	DB          *sql.DB
	MusicClient *clients.MusicInfoClient
}

func NewHandlers(db *sql.DB, musicClient *clients.MusicInfoClient) *Handlers {
	return &Handlers{
		DB:          db,
		MusicClient: musicClient}
}
