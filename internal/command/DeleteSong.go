package command

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// DeleteSong godoc
// @Summary Delete a song
// @Description Delete a song from the library
// @Tags songs
// @Param id path int true "Song ID"
//
//	@Success 201  {string} string "ok"
//
// @Failure 400  {string} string "Internal Server Error"
// @Failure 500  {string} string "Internal Server Error"
// @Router /songs/{id} [delete]
func (h *Handlers) DeleteSong(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM songs WHERE id=$1"
	_, err = h.DB.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
