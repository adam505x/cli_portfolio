package pages

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"portfolio/tui/router"
	"portfolio/tui/theme"
)

type AchievementsPage struct {
	offset int
}

func (p AchievementsPage) Update(key string) router.Action {
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

func (p AchievementsPage) View() string {
	// ── Left column: achievements content ─────────────────────────────────────
	var leftLines []string
	push := func(s string) { leftLines = append(leftLines, s) }

	push(theme.Heading.Render("ACHIEVEMENTS"))
	push(theme.Dim.Render(strings.Repeat("-", 44)))
	push("")

	push(theme.Base.Render("  Dr Seamus Mc Dermott Scholarship"))
	push(theme.Subtle.Render("  1 of 4 students in Ireland, awarded for impact on the economy and"))
	push(theme.Subtle.Render("  society and potential as a future entrepreneur"))
	push("")

	push(theme.Base.Render("  Naughton Scholarship"))
	push(theme.Subtle.Render("  1 of 36 students chosen nationwide for exceptional STEM achievements"))
	push(theme.Subtle.Render("  and outstanding extra-curricular achievements"))
	push("")

	push(theme.Base.Render("  Academic Excellence Award"))
	push(theme.Subtle.Render("  Graduated top of a class of 200 students"))
	push("")

	push(theme.Base.Render("  Irish Maths Olympiad 2023"))
	push(theme.Subtle.Render("  Contender in the final round, placing top fifty nationwide"))
	push("")

	push(theme.Base.Render("  1st Place - Microsoft PM Competition"))
	push(theme.Subtle.Render("  Selected to represent my school. Developed a mobile app MVP,"))
	push(theme.Subtle.Render("  pitched it to judges, and won against schools from across the country"))
	push("")

	push(theme.Base.Render("  Top 1% - MTU Computational Thinking Competition"))
	push("")

	push(theme.Base.Render("  2nd Place - UCD Hackathon"))
	push(theme.Subtle.Render("  Pitched Doclink, a platform for doctors who want to relocate"))
	push(theme.Subtle.Render("  to Ireland from abroad"))
	push("")

	push(theme.Base.Render("  Irish CTF Team"))
	push(theme.Subtle.Render("  Member of the Irish Team for CTFs and Attack/Defense competitions."))
	push(theme.Subtle.Render("  Shortlisted to compete in the European Cybersecurity Competition"))
	push(theme.Subtle.Render("  in Poland."))
	push("")

	push(theme.Dim.Render("esc  back"))

	// ── Right column: spinning earth ──────────────────────────────────────────
	var rightLines []string
	if len(EarthFrames) > 0 {
		frame := EarthFrames[AnimFrame%len(EarthFrames)]
		rightLines = strings.Split(frame, "\n")
	}

	maxLeft := 0
	for _, l := range leftLines {
		if w := lipgloss.Width(l); w > maxLeft {
			maxLeft = w
		}
	}

	const gap = "     "
	n := len(leftLines)
	if len(rightLines) > n {
		n = len(rightLines)
	}

	var b strings.Builder
	for i := 0; i < n; i++ {
		var left string
		if i < len(leftLines) {
			left = leftLines[i]
		}
		pad := strings.Repeat(" ", maxLeft-lipgloss.Width(left))
		var right string
		if i < len(rightLines) {
			right = rightLines[i]
		}
		b.WriteString(left + pad + gap + right + "\n")
	}

	return theme.Page.Render(clip(b.String(), p.offset, TermHeight-4))
}
