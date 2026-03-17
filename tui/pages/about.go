package pages

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"portfolio/tui/router"
	"portfolio/tui/services"
	"portfolio/tui/theme"
)

// TopTracks, RecentFilms, and AnimFrame are set by app.go.
var (
	TopTracks   []services.Track
	RecentFilms []services.Film
	AnimFrame   int
)

// ── Edit your personal info here ──────────────────────────────────────────────

var about = struct {
	Name     string
	Bio      []string
	GitHub   string
	LinkedIn string
	Email    string
}{
	Name:     "adam mcintyre",
	Bio:      []string{"building >"},
	GitHub:   "https://github.com/adam505x",
	LinkedIn: "https://www.linkedin.com/in/adam-mci/",
	Email:    "adam.mcintyre22@gmail.com",
}

// ─────────────────────────────────────────────────────────────────────────────

type AboutPage struct {
	offset int
}

func (p AboutPage) Update(key string) router.Action {
	switch key {
	case "esc":
		return router.Pop{}
	case "j", "down":
		p.offset++
		return router.Stay{Page: p}
	case "k", "up":
		if p.offset > 0 {
			p.offset--
		}
		return router.Stay{Page: p}
	}
	return router.Stay{Page: p}
}

func (p AboutPage) View() string {
	// ── Left column: horse animation ──────────────────────────────────────────
	var leftLines []string
	if len(HorseFrames) > 0 {
		frame := HorseFrames[AnimFrame%len(HorseFrames)]
		leftLines = strings.Split(frame, "\n")
	}

	// ── Right column: about content ───────────────────────────────────────────
	var right []string

	add := func(s string) { right = append(right, s) }
	row := func(label, value string) string {
		return fmt.Sprintf("%s   %s",
			theme.Subtle.Width(12).Render(label),
			theme.Base.Render(value),
		)
	}

	add(theme.Heading.Render("ABOUT"))
	add("")
	add(row("name", about.Name))
	add("")
	for i, line := range about.Bio {
		label := ""
		if i == 0 {
			label = "bio"
		}
		add(row(label, line))
	}
	add("")
	add(row("github", about.GitHub))
	add(row("linkedin", about.LinkedIn))
	add(row("email", about.Email))

	add("")
	add(theme.Dim.Render(strings.Repeat("─", 36)))
	add("")
	add(theme.Heading.Render("RECENTLY WATCHED"))
	add("")

	if len(RecentFilms) == 0 {
		add(theme.Subtle.Render("  loading..."))
	} else {
		for _, f := range RecentFilms {
			title := fmt.Sprintf("%s (%s)", f.Name, f.Year)
			add(fmt.Sprintf("%s   %s",
				theme.Base.Width(30).Render(title),
				theme.Subtle.Render(f.Stars()),
			))
		}
	}

	add("")
	add(theme.Dim.Render(strings.Repeat("─", 36)))
	add("")
	add(theme.Heading.Render("TOP TRACKS THIS MONTH"))
	add("")

	if len(TopTracks) == 0 {
		add(theme.Subtle.Render("  loading..."))
	} else if TopTracks[0].Name == "—" {
		add(theme.Subtle.Render("  unavailable — set Last.fm credentials in .env"))
	} else {
		for i, t := range TopTracks {
			add(theme.Base.Render(fmt.Sprintf("  %d. %s — %s", i+1, t.Name, t.Artist)))
			add(theme.Subtle.Render(fmt.Sprintf("     %s plays", t.PlayCount)))
		}
	}

	add("")
	add(theme.Dim.Render("esc  back"))

	// ── Manual line-by-line horizontal join (about left, horse right) ──────────
	maxLeft := 0
	for _, l := range right {
		if w := lipgloss.Width(l); w > maxLeft {
			maxLeft = w
		}
	}
	const gap = "          " // padding between columns

	n := len(right)
	if len(leftLines) > n {
		n = len(leftLines)
	}

	var b strings.Builder
	for i := 0; i < n; i++ {
		var left string
		if i < len(right) {
			left = right[i]
		}
		pad := strings.Repeat(" ", maxLeft-lipgloss.Width(left))

		var r string
		if i < len(leftLines) {
			r = leftLines[i]
		}
		b.WriteString(left + pad + gap + r + "\n")
	}

	return theme.Page.Render(clip(b.String(), p.offset, TermHeight-4))
}
