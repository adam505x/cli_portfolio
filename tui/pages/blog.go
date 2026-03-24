package pages

import (
	"strings"

	"portfolio/tui/router"
	"portfolio/tui/theme"
)

type BlogPage struct {
	offset int
}

func (p BlogPage) Update(key string) router.Action {
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

func (p BlogPage) View() string {
	var lines []string
	add := func(s string) { lines = append(lines, s) }

	add(theme.Heading.Render("BLOG"))
	add(theme.Dim.Render(strings.Repeat("-", 40)))
	add("")
	add(theme.Subtle.Render("coming soon"))
	add("")

	if GalaxyArt != "" {
		for _, l := range strings.Split(strings.TrimRight(GalaxyArt, "\n\r"), "\n") {
			add(theme.Base.Render(l))
		}
	}

	add("")
	add(theme.Dim.Render("esc  back"))

	return theme.Page.Render(clip(strings.Join(lines, "\n"), p.offset, TermHeight-4))
}
