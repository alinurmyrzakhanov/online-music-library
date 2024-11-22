package command

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

// UpdateSong godoc
// @Summary Update a song
// @Description Update details of an existing song
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "Song ID"
// @Param song body Song true "Updated song data"
// @Success 201  {string} string "ok"
// @Failure 400  {string} string "Internal Server Error"
// @Failure 500  {string} string "Internal Server Error"
// @Router /songs/{id} [put]
func (h *Handlers) UpdateSong(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedSong Song
	if err := json.NewDecoder(r.Body).Decode(&updatedSong); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "UPDATE songs SET group_name=$1, song_name=$2, release_date=$3, text=$4, link=$5, updated_at=$6 WHERE id=$7"
	_, err = h.DB.Exec(query, updatedSong.Group, updatedSong.Song, updatedSong.ReleaseDate, updatedSong.Text, updatedSong.Link, time.Now(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
