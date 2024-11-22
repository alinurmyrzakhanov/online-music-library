package command

import (
	"encoding/json"
	"net/http"
)

// GetSongs godoc
// @Summary Get all songs
// @Description Get details of all songs
// @Tags songs
// @Produce json
// @Success 200  {string} string "ok"
// Failure 500  {string} string "Internal Server Error"
// @Router /songs [get]
func (h *Handlers) GetSongs(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT id, group_name, song_name, release_date, text, link FROM songs")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var songs []Song
	for rows.Next() {
		var song Song
		if err := rows.Scan(&song.ID, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		songs = append(songs, song)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songs)
}
