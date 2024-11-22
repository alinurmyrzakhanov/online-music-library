package command

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// GetSongByID godoc
// @Summary Get song by ID
// @Description Get details of a song by ID
// @Tags songs
// @Param id path int true "Song ID"
// @Produce json
//
//	@Success 201  {string} string "ok"
//
// @Failure 400  {string} string "Internal Server Error"
// @Failure 500  {string} string "Internal Server Error"
// @Router /songs/{id} [get]
func (h *Handlers) GetSongByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var song Song
	query := "SELECT id, group_name, song_name, release_date, text, link FROM songs WHERE id = $1"
	row := h.DB.QueryRow(query, id)
	if err := row.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Song not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(song)
}
