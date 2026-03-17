package pages

import (
	"strings"

	"portfolio/tui/router"
	"portfolio/tui/theme"
)

// ── Edit your projects here ───────────────────────────────────────────────────

var projects = []ProjectData{
	{
		Name: "cli-portfolio",
		Desc: "This terminal portfolio, served over SSH.",
		Link: "github.com/you/cli-portfolio",
		Body: "Built with Go and the Charm.sh ecosystem (bubbletea, lipgloss, wish).\nServed over SSH on port 2222 and over the web via WebSocket.",
	},
	{
		Name: "gobox",
		Desc: "A minimal container runtime written in Go.",
		Link: "github.com/you/gobox",
		Body: "Implements Linux namespaces and cgroups from scratch.\nNo Docker dependency — just syscalls and Go.",
	},
	{
		Name: "tinydb",
		Desc: "An embeddable key-value store for Go apps.",
		Link: "github.com/you/tinydb",
		Body: "LSM-tree based storage engine with WAL for crash recovery.\nDesigned to be embedded, not run as a server.",
	},
}

// ── ProjectData is what you fill in per project ───────────────────────────────

type ProjectData struct {
	Name string // short name, shown in the list
	Desc string // one-line description
	Link string // github or url
	Body string // full detail shown on the project's own page (newlines ok)
}

// ─────────────────────────────────────────────────────────────────────────────

type ProjectsPage struct {
	cursor int
}

func (p ProjectsPage) Update(key string) router.Action {
	switch key {
	case "esc":
		return router.Pop{}
	case "enter":
		return router.Push{Page: NewProjectPage(projects[p.cursor])}
	case "j", "down":
		if p.cursor < len(projects)-1 {
			p.cursor++
		}
		return router.Stay{Page: p}
	case "k", "up":
		if p.cursor > 0 {
			p.cursor--
		}
		return router.Stay{Page: p}
	}
	return router.Stay{Page: p}
}

func (p ProjectsPage) View() string {
	var b strings.Builder

	b.WriteString(theme.Heading.Render("PROJECTS"))
	b.WriteString("\n")
	b.WriteString(theme.Dim.Render(strings.Repeat("─", 40)))
	b.WriteString("\n\n")

	for i, proj := range projects {
		cursor := "  "
		if i == p.cursor {
			cursor = "> "
		}
		b.WriteString(theme.Accent.Render(cursor + proj.Name))
		b.WriteString("\n")
		b.WriteString(theme.Base.Render("    " + proj.Desc))
		b.WriteString("\n")
		b.WriteString(theme.Subtle.Render("    " + proj.Link))
		b.WriteString("\n\n")
	}

	b.WriteString("\n")
	b.WriteString(theme.Dim.Render("j/k  move   enter  open   esc  back"))

	return theme.Page.Render(b.String())
}
