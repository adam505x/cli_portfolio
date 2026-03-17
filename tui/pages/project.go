package pages

import (
	"strings"

	"portfolio/tui/router"
	"portfolio/tui/theme"
)

// ProjectPage is the detail view for a single project.
// It is pushed onto the stack when the user presses enter on the projects list.
type ProjectPage struct {
	data ProjectData
}

func NewProjectPage(d ProjectData) ProjectPage {
	return ProjectPage{data: d}
}

func (p ProjectPage) Update(key string) router.Action {
	switch key {
	case "esc", "q":
		return router.Pop{}
	}
	return router.Stay{Page: p}
}

func (p ProjectPage) View() string {
	var b strings.Builder

	b.WriteString(theme.Heading.Render(p.data.Name))
	b.WriteString("\n")
	b.WriteString(theme.Dim.Render(strings.Repeat("─", 40)))
	b.WriteString("\n\n")

	for _, line := range strings.Split(p.data.Body, "\n") {
		b.WriteString(theme.Base.Render(line))
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(theme.Subtle.Render(p.data.Link))
	b.WriteString("\n\n")
	b.WriteString(theme.Dim.Render("esc  back"))

	return theme.Page.Render(b.String())
}
