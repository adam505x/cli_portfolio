package pages

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed frames_earth
var earthFS embed.FS

// EarthFrames holds all earth animation frames loaded at startup.
var EarthFrames []string

func init() {
	for i := 1; i <= 44; i++ {
		name := fmt.Sprintf("frames_earth/frame%02d.txt", i)
		data, err := earthFS.ReadFile(name)
		if err != nil {
			EarthFrames = append(EarthFrames, "")
			continue
		}
		lines := strings.Split(strings.TrimRight(string(data), "\n\r"), "\n")
		if len(lines) > 13 {
			lines = lines[3 : len(lines)-10]
		}
		EarthFrames = append(EarthFrames, strings.Join(lines, "\n"))
	}
}
