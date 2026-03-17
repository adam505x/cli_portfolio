package pages

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed frames
var framesFS embed.FS

// HorseFrames holds all animation frames loaded at startup.
var HorseFrames []string

func init() {
	for i := 1; i <= 12; i++ {
		name := fmt.Sprintf("frames/frame%02d.txt", i)
		data, err := framesFS.ReadFile(name)
		if err != nil {
			HorseFrames = append(HorseFrames, "")
			continue
		}
		// Strip trailing newline so frames are consistent
		HorseFrames = append(HorseFrames, strings.TrimRight(string(data), "\n"))
	}
}
