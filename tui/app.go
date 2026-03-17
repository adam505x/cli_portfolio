package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type screen int

const (
	homeScreen screen = iota
	projectsScreen
	blogScreen
	aboutScreen
)

// ── Data ────────────────────────────────────────────────────────────────────

type project struct {
	name, desc, link string
}

type post struct {
	title, date string
}

var projects = []project{
	{"cli-portfolio", "This terminal portfolio, served over SSH.", "github.com/you/cli-portfolio"},
	{"gobox", "A minimal container runtime written in Go.", "github.com/you/gobox"},
	{"tinydb", "An embeddable key-value store for Go apps.", "github.com/you/tinydb"},
}

var posts = []post{
	{"Building a Terminal Portfolio with Charm.sh", "2026-03-01"},
	{"Go Concurrency Patterns I Actually Use",      "2026-01-14"},
	{"Why I Switched Back to the Terminal",         "2025-11-22"},
}

// ── Model ───────────────────────────────────────────────────────────────────

type Model struct {
	current screen
	width   int
	height  int
}

func NewModel(w, h int) Model {
	return Model{current: homeScreen, width: w, height: h}
}

// ── Init / Update ────────────────────────────────────────────────────────────

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "1":
			m.current = projectsScreen
		case "2":
			m.current = blogScreen
		case "3":
			m.current = aboutScreen
		case "esc":
			m.current = homeScreen
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

// ── View ────────────────────────────────────────────────────────────────────

func (m Model) View() string {
	bg := lipgloss.NewStyle().
		Background(lipgloss.Color(colorBg)).
		Width(m.width).
		Height(m.height)

	var content string
	switch m.current {
	case homeScreen:
		content = m.viewHome()
	case projectsScreen:
		content = m.viewProjects()
	case blogScreen:
		content = m.viewBlog()
	case aboutScreen:
		content = m.viewAbout()
	}

	return bg.Render(content)
}

// ── Home ─────────────────────────────────────────────────────────────────────

func (m Model) viewHome() string {
	// Split portrait into lines; pad each to the same rune-width so every
	// line is the same column-count regardless of how the terminal renders
	// the characters (avoids lipgloss.JoinHorizontal measurement issues).
	portraitLines := strings.Split(strings.TrimRight(Portrait, "\n"), "\n")
	maxRunes := 0
	for _, l := range portraitLines {
		if w := lipgloss.Width(l); w > maxRunes {
			maxRunes = w
		}
	}

	// Build right-column lines as plain strings (no ANSI yet) so we can
	// concatenate them with portrait lines before applying any style.
	nameLines := strings.Split(strings.TrimRight(nameArtRaw, "\n"), "\n")
	rightLines := make([]string, 0, len(nameLines)+14)
	rightLines = append(rightLines, nameLines...)
	rightLines = append(rightLines,
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

	const colGap = "    " // 4-space gap between portrait and content
	nLines := len(portraitLines)
	if len(rightLines) > nLines {
		nLines = len(rightLines)
	}

	var b strings.Builder
	for i := 0; i < nLines; i++ {
		var portLine string
		if i < len(portraitLines) {
			portLine = portraitLines[i]
		}
		// Pad portrait segment to maxRunes so every row is the same width
		pad := strings.Repeat(" ", maxRunes-lipgloss.Width(portLine))

		var rightLine string
		if i < len(rightLines) {
			rightLine = rightLines[i]
		}

		b.WriteString(portLine + pad + colGap + rightLine + "\n")
	}

	return baseStyle.Render(b.String())
}

func navHint(key, label string) string {
	return fmt.Sprintf("%s  %s",
		keyStyle.Render("press "+key),
		baseStyle.Render("to see my "+label),
	)
}

// ── Projects ─────────────────────────────────────────────────────────────────

func (m Model) viewProjects() string {
	var b strings.Builder

	b.WriteString(headingStyle.Render("PROJECTS"))
	b.WriteString("\n")
	b.WriteString(dimStyle.Render(strings.Repeat("─", 40)))
	b.WriteString("\n\n")

	for _, p := range projects {
		b.WriteString(accentStyle.Render(p.name))
		b.WriteString("\n")
		b.WriteString(baseStyle.Render("  "+p.desc))
		b.WriteString("\n")
		b.WriteString(subtleStyle.Render("  "+p.link))
		b.WriteString("\n\n")
	}

	b.WriteString("\n")
	b.WriteString(dimStyle.Render("esc  back"))

	return pageStyle.Render(b.String())
}

// ── Blog ─────────────────────────────────────────────────────────────────────

func (m Model) viewBlog() string {
	var b strings.Builder

	b.WriteString(headingStyle.Render("BLOG POSTS"))
	b.WriteString("\n")
	b.WriteString(dimStyle.Render(strings.Repeat("─", 40)))
	b.WriteString("\n\n")

	for _, p := range posts {
		b.WriteString(accentStyle.Render(p.title))
		b.WriteString("\n")
		b.WriteString(subtleStyle.Render("  "+p.date))
		b.WriteString("\n\n")
	}

	b.WriteString("\n")
	b.WriteString(dimStyle.Render("esc  back"))

	return pageStyle.Render(b.String())
}

// ── About ────────────────────────────────────────────────────────────────────

func (m Model) viewAbout() string {
	var b strings.Builder

	b.WriteString(headingStyle.Render("ABOUT"))
	b.WriteString("\n")
	b.WriteString(dimStyle.Render(strings.Repeat("─", 40)))
	b.WriteString("\n\n")

	row := func(label, value string) string {
		return fmt.Sprintf("%s   %s\n",
			subtleStyle.Width(10).Render(label),
			baseStyle.Render(value),
		)
	}

	b.WriteString(row("name", "Adam McIntyre"))
	b.WriteString("\n")
	b.WriteString(row("bio", "Software engineer. Interested in systems,"))
	b.WriteString(row("", "    tooling, and terminal UIs."))
	b.WriteString("\n")
	b.WriteString(row("github", "github.com/you"))
	b.WriteString(row("linkedin", "linkedin.com/in/you"))
	b.WriteString(row("email", "hello@example.com"))

	b.WriteString("\n\n")
	b.WriteString(dimStyle.Render("esc  back"))

	return pageStyle.Render(b.String())
}
