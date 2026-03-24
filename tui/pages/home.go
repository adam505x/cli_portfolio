package pages

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"portfolio/tui/router"
	"portfolio/tui/theme"
)

var homeMenuItems = []struct {
	label string
	page  func() router.Page
}{
	{"projects", func() router.Page { return ProjectsPage{} }},
	{"blog", func() router.Page { return BlogPage{} }},
	{"about me", func() router.Page { return AboutPage{} }},
	{"achievements", func() router.Page { return AchievementsPage{} }},
	{"guest book", func() router.Page { return NewGuestbookPage() }},
}

type HomePage struct {
	portrait string
	nameArt  string
	cursor   int
}

func NewHome(portrait, nameArt string) HomePage {
	return HomePage{portrait: portrait, nameArt: nameArt}
}

func (h HomePage) Update(key string) router.Action {
	switch key {
	case "q":
		return router.Quit{}
	case "up", "k":
		if h.cursor > 0 {
			h.cursor--
		}
		return router.Stay{Page: h}
	case "down", "j":
		if h.cursor < len(homeMenuItems)-1 {
			h.cursor++
		}
		return router.Stay{Page: h}
	case "enter":
		return router.Push{Page: homeMenuItems[h.cursor].page()}
	}
	return router.Stay{Page: h}
}

func (h HomePage) View() string {
	portraitLines := strings.Split(strings.TrimRight(h.portrait, "\n"), "\n")
	maxWidth := 0
	for _, l := range portraitLines {
		if w := lipgloss.Width(l); w > maxWidth {
			maxWidth = w
		}
	}

	nameLines := strings.Split(strings.TrimRight(h.nameArt, "\n"), "\n")
	right := make([]string, 0, len(nameLines)+16)
	right = append(right, nameLines...)
	right = append(right, "", "", "welcome to my website", "(use arrow keys to navigate)", "")

	for i, item := range homeMenuItems {
		cursor := "  "
		if i == h.cursor {
			cursor = "> "
		}
		right = append(right, cursor+item.label)
	}

	right = append(right, "", "", "up/down  move   enter  open   q  quit")

	const gap = "    "
	n := len(portraitLines)
	if len(right) > n {
		n = len(right)
	}

	var b strings.Builder
	for i := 0; i < n; i++ {
		var left string
		if i < len(portraitLines) {
			left = portraitLines[i]
		}
		pad := strings.Repeat(" ", maxWidth-lipgloss.Width(left))

		var r string
		if i < len(right) {
			r = right[i]
		}
		b.WriteString(left + pad + gap + r + "\n")
	}

	return theme.Base.Render(b.String())
}
