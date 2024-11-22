package command

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// AddSong godoc
// @Summary Add a new song
// @Description Add a new song to the library
// @Tags songs
// @Accept json
// @Produce json
// @Param song body Song true "Song data"
// @Success 201  {string} string "ok"
// @Failure 400  {string} string "Internal Server Error"
// @Failure 500  {string} string "Internal Server Error"
// @Router /songs [post]
func (h *Handlers) AddSong(w http.ResponseWriter, r *http.Request) {
	var song Song
	if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	songDetail, err := h.MusicClient.GetSongInfo(song.Group, song.Song)
	if err != nil {
		log.Printf("Failed to get song info from external API: %v", err)
		log.Println("Proceeding with partial song data")
	} else {
		song.ReleaseDate = songDetail.ReleaseDate
		song.Text = songDetail.Text
		song.Link = songDetail.Link
		log.Printf("Song info enriched: ReleaseDate=%s, Text=%s, Link=%s",
			songDetail.ReleaseDate, songDetail.Text, songDetail.Link)
	}

	query := "INSERT INTO songs (group_name, song_name, release_date, text, link, created_at) VALUES ($1, $2, $3, $4, $5)"
	_, err = h.DB.Exec(query, song.Group, song.Song, song.ReleaseDate, song.Text, song.Link, time.Now())
	if err != nil {
		log.Printf("Failed to save song in the database: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("[INFO] Song added successfully: Group=%s, Song=%s", song.Group, song.Song)

	w.WriteHeader(http.StatusCreated)
}
