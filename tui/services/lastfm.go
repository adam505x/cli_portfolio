package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Track struct {
	Name      string
	Artist    string
	PlayCount string
}

func FetchTopTracks() ([]Track, error) {
	key, user := lastFMKey(), lastFMUser()
	if key == "" || key == "your_api_key_here" {
		return nil, fmt.Errorf("LASTFM_KEY not set")
	}

	url := fmt.Sprintf(
		"https://ws.audioscrobbler.com/2.0/?method=user.getTopTracks&user=%s&period=1month&limit=3&api_key=%s&format=json",
		user, key,
	)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		TopTracks struct {
			Track []struct {
				Name      string `json:"name"`
				PlayCount string `json:"playcount"`
				Artist    struct {
					Name string `json:"name"`
				} `json:"artist"`
			} `json:"track"`
		} `json:"toptracks"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if len(result.TopTracks.Track) == 0 {
		return nil, fmt.Errorf("no tracks found")
	}

	var tracks []Track
	for _, t := range result.TopTracks.Track {
		tracks = append(tracks, Track{Name: t.Name, Artist: t.Artist.Name, PlayCount: t.PlayCount})
	}
	return tracks, nil
}
