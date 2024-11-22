package clients

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type MusicInfoClient struct {
	BaseURL string
}

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func NewMusicInfoClient(baseURL string) *MusicInfoClient {
	return &MusicInfoClient{BaseURL: baseURL}
}

func (c *MusicInfoClient) GetSongInfo(group, song string) (*SongDetail, error) {
	log.Printf("Fetching song info for group: %s, song: %s", group, song)
	endpoint := fmt.Sprintf("%s/info", c.BaseURL)
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("group", group)
	q.Set("song", song)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Printf("HTTP request failed: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected response status: %d", resp.StatusCode)
		return nil, fmt.Errorf("failed to get song info: status code %d", resp.StatusCode)
	}

	var songDetail SongDetail
	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		log.Printf("Failed to decode response body: %v", err)
		return nil, err
	}

	return &songDetail, nil
}
