package pages

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
	"portfolio/tui/router"
	"portfolio/tui/theme"
)

const guestbookFile = "guestbook.json"

type gbEntry struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Time    string `json:"time"`
}

type GuestbookPage struct {
	entries []gbEntry
	mode    string // "view" | "name" | "message"
	nameBuf string
	msgBuf  string
	offset  int
}

func NewGuestbookPage() GuestbookPage {
	return GuestbookPage{entries: loadGuestbook(), mode: "view"}
}

func loadGuestbook() []gbEntry {
	data, err := os.ReadFile(guestbookFile)
	if err != nil {
		return nil
	}
	var entries []gbEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil
	}
	return entries
}

func saveGuestbook(entries []gbEntry) {
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return
	}
	os.WriteFile(guestbookFile, data, 0644) //nolint:errcheck
}

func (p GuestbookPage) Update(key string) router.Action {
	switch p.mode {
	case "view":
		switch key {
		case "esc":
			return router.Pop{}
		case "n":
			p.mode = "name"
			p.nameBuf = ""
			p.msgBuf = ""
			return router.Stay{Page: p}
		case "j", "down":
			p.offset++
			return router.Stay{Page: p}
		case "k", "up":
			if p.offset > 0 {
				p.offset--
			}
			return router.Stay{Page: p}
		}

	case "name":
		switch key {
		case "esc":
			p.mode = "view"
			return router.Stay{Page: p}
		case "enter":
			if strings.TrimSpace(p.nameBuf) != "" {
				p.mode = "message"
			}
			return router.Stay{Page: p}
		case "backspace":
			runes := []rune(p.nameBuf)
			if len(runes) > 0 {
				p.nameBuf = string(runes[:len(runes)-1])
			}
			return router.Stay{Page: p}
		default:
			if utf8.RuneCountInString(key) == 1 && len([]rune(p.nameBuf)) < 30 {
				p.nameBuf += key
			}
			return router.Stay{Page: p}
		}

	case "message":
		switch key {
		case "esc":
			p.mode = "view"
			return router.Stay{Page: p}
		case "enter":
			if strings.TrimSpace(p.msgBuf) != "" {
				entry := gbEntry{
					Name:    strings.TrimSpace(p.nameBuf),
					Message: strings.TrimSpace(p.msgBuf),
					Time:    time.Now().Format("2006-01-02 15:04"),
				}
				p.entries = append(p.entries, entry)
				saveGuestbook(p.entries)
				p.mode = "view"
				p.nameBuf = ""
				p.msgBuf = ""
			}
			return router.Stay{Page: p}
		case "backspace":
			runes := []rune(p.msgBuf)
			if len(runes) > 0 {
				p.msgBuf = string(runes[:len(runes)-1])
			}
			return router.Stay{Page: p}
		default:
			if utf8.RuneCountInString(key) == 1 && len([]rune(p.msgBuf)) < 100 {
				p.msgBuf += key
			}
			return router.Stay{Page: p}
		}
	}
	return router.Stay{Page: p}
}

func (p GuestbookPage) View() string {
	var leftLines []string
	push := func(s string) { leftLines = append(leftLines, s) }

	push(theme.Heading.Render("GUEST BOOK"))
	push(theme.Dim.Render(strings.Repeat("-", 44)))
	push("")

	if p.mode == "name" || p.mode == "message" {
		push(theme.Subtle.Render("leave a message"))
		push("")

		nameCursor := ""
		msgCursor := ""
		if p.mode == "name" {
			nameCursor = "█"
		} else {
			msgCursor = "█"
		}

		push(theme.Base.Render(fmt.Sprintf("  name      %s%s", p.nameBuf, nameCursor)))
		push(theme.Base.Render(fmt.Sprintf("  message   %s%s", p.msgBuf, msgCursor)))
		push("")
		push(theme.Dim.Render("enter  confirm   esc  cancel"))

		return theme.Page.Render(strings.Join(leftLines, "\n"))
	}

	if len(p.entries) == 0 {
		push(theme.Subtle.Render("  no messages yet - be the first!"))
	} else {
		for _, e := range p.entries {
			nameRunes := []rune(e.Name)
			name := e.Name
			if len(nameRunes) > 20 {
				name = string(nameRunes[:20])
			}
			namePad := name + strings.Repeat(" ", 20-len([]rune(name)))
			push(fmt.Sprintf("  %s   %s   %s",
				theme.Accent.Render(namePad),
				theme.Dim.Render(e.Time),
				theme.Base.Render(e.Message),
			))
			push("")
		}
	}

	push("")
	push(theme.Dim.Render("n  sign   esc  back"))

	// ── Right column: cat art ─────────────────────────────────────────────────
	var rightLines []string
	if CatArt != "" {
		rightLines = strings.Split(strings.TrimRight(CatArt, "\n\r"), "\n")
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
