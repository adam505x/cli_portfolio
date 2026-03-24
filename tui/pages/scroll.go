package pages

import "strings"

// TermHeight and TermWidth are set by app.go whenever the terminal is resized.
var (
	TermHeight int
	TermWidth  int
)

// contentWidth returns the usable width inside theme.Page's Padding(2,4).
func contentWidth() int {
	w := TermWidth - 8 // 4 left + 4 right padding
	if w < 40 {
		return 80 // sane default before first resize
	}
	return w
}

// wordWrap breaks s into lines no longer than width, splitting at word boundaries.
func wordWrap(s string, width int) string {
	if width <= 0 || len(s) == 0 {
		return s
	}
	words := strings.Fields(s)
	if len(words) == 0 {
		return s
	}
	var lines []string
	var cur strings.Builder
	curLen := 0
	for _, w := range words {
		wl := len([]rune(w))
		if curLen == 0 {
			cur.WriteString(w)
			curLen = wl
		} else if curLen+1+wl <= width {
			cur.WriteByte(' ')
			cur.WriteString(w)
			curLen += 1 + wl
		} else {
			lines = append(lines, cur.String())
			cur.Reset()
			cur.WriteString(w)
			curLen = wl
		}
	}
	if cur.Len() > 0 {
		lines = append(lines, cur.String())
	}
	return strings.Join(lines, "\n")
}

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
