package services

import (
	"encoding/csv"
	"math"
	"strconv"
	"strings"
)

type Film struct {
	Name        string
	Year        string
	Rating      float64
	WatchedDate string
}

// Stars converts a numeric rating (0.5–5.0) to a unicode star string.
func (f Film) Stars() string {
	full := int(math.Floor(f.Rating))
	half := f.Rating-float64(full) >= 0.5
	var s strings.Builder
	for i := 0; i < full; i++ {
		s.WriteRune('★')
	}
	if half {
		s.WriteString("½")
	}
	return s.String()
}

// ParseDiary parses a Letterboxd diary CSV export and returns the most
// recent n films, newest first.
func ParseDiary(csv string, n int) ([]Film, error) {
	films, err := parseCSV(csv)
	if err != nil {
		return nil, err
	}
	// CSV is oldest-first; reverse to get newest first
	for i, j := 0, len(films)-1; i < j; i, j = i+1, j-1 {
		films[i], films[j] = films[j], films[i]
	}
	if len(films) > n {
		films = films[:n]
	}
	return films, nil
}

func parseCSV(data string) ([]Film, error) {
	r := csv.NewReader(strings.NewReader(data))
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	var films []Film
	for _, row := range records[1:] { // skip header
		if len(row) < 6 {
			continue
		}
		rating, _ := strconv.ParseFloat(row[4], 64)
		films = append(films, Film{
			Name:        row[1],
			Year:        row[2],
			Rating:      rating,
			WatchedDate: row[7],
		})
	}
	return films, nil
}
