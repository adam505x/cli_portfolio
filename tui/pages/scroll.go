package pages

import "strings"

// TermHeight is set by app.go whenever the terminal is resized.
var TermHeight int

// clip returns the lines of content visible given an offset and available height.
// It also appends a subtle scroll indicator when more content exists above/below.
func clip(content string, offset, height int) string {
	if height <= 0 {
		return content
	}
	lines := strings.Split(content, "\n")

	// clamp offset
	max := len(lines) - height
	if max < 0 {
		max = 0
	}
	if offset > max {
		offset = max
	}
	if offset < 0 {
		offset = 0
	}

	end := offset + height
	if end > len(lines) {
		end = len(lines)
	}

	visible := lines[offset:end]

	// scroll indicators on the last visible line
	indicator := ""
	if offset > 0 && offset+height < len(lines) {
		indicator = "↑↓ scroll"
	} else if offset > 0 {
		indicator = "↑ scroll"
	} else if offset+height < len(lines) {
		indicator = "↓ scroll"
	}
	if indicator != "" {
		if len(visible) > 0 {
			visible[len(visible)-1] = "  " + indicator
		}
	}

	return strings.Join(visible, "\n")
}
