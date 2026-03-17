package pages

import (
	"strings"

	"portfolio/tui/router"
	"portfolio/tui/theme"
)

// ── Edit your blog posts here ─────────────────────────────────────────────────

var posts = []struct {
	Title, Date string
}{
	{"Building a Terminal Portfolio with Charm.sh", "2026-03-01"},
	{"Go Concurrency Patterns I Actually Use", "2026-01-14"},
	{"Why I Switched Back to the Terminal", "2025-11-22"},
}

// ─────────────────────────────────────────────────────────────────────────────

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
	var b strings.Builder

	b.WriteString(theme.Heading.Render("BLOG POSTS"))
	b.WriteString("\n")
	b.WriteString(theme.Dim.Render(strings.Repeat("─", 40)))
	b.WriteString("\n\n")

	for _, post := range posts {
		b.WriteString(theme.Accent.Render(post.Title))
		b.WriteString("\n")
		b.WriteString(theme.Subtle.Render("  " + post.Date))
		b.WriteString("\n\n")
	}

	b.WriteString("\n")
	b.WriteString(theme.Dim.Render("esc  back"))

	return theme.Page.Render(clip(b.String(), p.offset, TermHeight-4))
}
