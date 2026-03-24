package pages

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed frames_horse
var frames_horseFS embed.FS

// HorseFrames holds all animation frames_horse loaded at startup.
var HorseFrames []string

func init() {
	for i := 1; i <= 12; i++ {
		name := fmt.Sprintf("frames_horse/frame%02d.txt", i)
		data, err := frames_horseFS.ReadFile(name)
		if err != nil {
			HorseFrames = append(HorseFrames, "")
			continue
		}
		// Strip trailing newline so frames_horse are consistent
		HorseFrames = append(HorseFrames, strings.TrimRight(string(data), "\n"))
	}
}
