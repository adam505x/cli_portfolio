package pages

import (
	"fmt"
	"strings"

	"portfolio/tui/router"
	"portfolio/tui/theme"
)

// ProjectPage is the detail view for a single project.
type ProjectPage struct {
	data   ProjectData
	offset int
}

func NewProjectPage(d ProjectData) ProjectPage {
	return ProjectPage{data: d}
}

func (p ProjectPage) Update(key string) router.Action {
	switch key {
	case "esc", "q":
		return router.Pop{}
	case "down":
		p.offset++
		return router.Stay{Page: p}
	case "up":
		if p.offset > 0 {
			p.offset--
		}
		return router.Stay{Page: p}
	}
	return router.Stay{Page: p}
}

// osc8 wraps text in an OSC 8 terminal hyperlink.
func osc8(url, text string) string {
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}
	return "\x1b]8;;" + url + "\x1b\\" + text + "\x1b]8;;\x1b\\"
}

func (p ProjectPage) View() string {
	var lines []string
	add := func(s string) { lines = append(lines, s) }

	row := func(label, value string) string {
		return fmt.Sprintf("%s   %s",
			theme.Subtle.Width(12).Render(label),
			theme.Base.Render(value),
		)
	}

	// ── Title ─────────────────────────────────────────────────────────────────
	add(theme.Heading.Render(p.data.Name))
	add(theme.Dim.Render(strings.Repeat("─", 44)))
	add("")

	// ── Body ──────────────────────────────────────────────────────────────────
	wrap := contentWidth()
	for _, line := range strings.Split(strings.TrimSpace(p.data.Body), "\n") {
		if strings.HasPrefix(line, "## ") {
			add("")
			add(theme.Heading.Render(strings.TrimPrefix(line, "## ")))
			add(theme.Dim.Render(strings.Repeat("─", 24)))
		} else if line == "" {
			add("")
		} else {
			for _, wrapped := range strings.Split(wordWrap(line, wrap), "\n") {
				add(theme.Base.Render(wrapped))
			}
		}
	}

	// ── Reference quote ───────────────────────────────────────────────────────
	if p.data.Quote != "" {
		add("")
		add(theme.Heading.Render("REFERENCE"))
		add(theme.Dim.Render(strings.Repeat("-", 24)))
		add("")
		add(theme.Dim.Render(` _  _`))
		add(theme.Dim.Render(`| || |`))
		add(theme.Dim.Render(`\_|\_|`))
		add("")
		indentWidth := contentWidth() - 6
		for _, line := range strings.Split(p.data.Quote, "\n") {
			if strings.TrimSpace(line) == "" {
				add("")
			} else {
				for _, wrapped := range strings.Split(wordWrap(strings.TrimSpace(line), indentWidth), "\n") {
					add(theme.Base.Render("      " + wrapped))
				}
			}
		}
		add("")
		cw := contentWidth()
		for _, cl := range []string{` _  _`, `| || |`, `\_|\_|`} {
			runeLen := len([]rune(cl))
			pad := ""
			if cw > runeLen {
				pad = strings.Repeat(" ", cw-runeLen)
			}
			add(theme.Dim.Render(pad + cl))
		}
		add("")
		add(theme.Subtle.Render("- " + p.data.QuoteBy))
	}

	// ── Tech stack ────────────────────────────────────────────────────────────
	if p.data.Tech != "" {
		add("")
		add(theme.Dim.Render(strings.Repeat("─", 44)))
		add(row("Tech Stack:", p.data.Tech))
	}

	// ── Links ─────────────────────────────────────────────────────────────────
	if len(p.data.Links) > 0 {
		add("")
		add(theme.Dim.Render(strings.Repeat("─", 44)))
		add(theme.Heading.Render("LINKS"))
		add("")
		for _, link := range p.data.Links {
			add("  " + osc8(link, link))
		}
	}

	add("")
	add(theme.Dim.Render(strings.Repeat("─", 44)))
	add(theme.Dim.Render("↑/↓  scroll   esc  back"))

	content := strings.Join(lines, "\n")
	return theme.Page.Render(clip(content, p.offset, TermHeight-4))
}
