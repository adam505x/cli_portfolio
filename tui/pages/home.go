package pages

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"portfolio/tui/router"
	"portfolio/tui/theme"
)

type HomePage struct {
	portrait string
	nameArt  string
}

func NewHome(portrait, nameArt string) HomePage {
	return HomePage{portrait: portrait, nameArt: nameArt}
}

func (h HomePage) Update(key string) router.Action {
	switch key {
	case "q":
		return router.Quit{}
	case "1":
		return router.Push{Page: ProjectsPage{}}
	case "2":
		return router.Push{Page: BlogPage{}}
	case "3":
		return router.Push{Page: AboutPage{}}
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
	right := make([]string, 0, len(nameLines)+12)
	right = append(right, nameLines...)
	right = append(right,
		"",
		"",
		"welcome to my portfolio",
		"",
		"press 1   to see my projects",
		"",
		"press 2   to see my blog posts",
		"",
		"press 3   to learn more about me",
		"",
		"",
		"q  quit",
	)

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
