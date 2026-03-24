package theme

import "github.com/charmbracelet/lipgloss"

const (
	ColorBg    = "#000000"
	ColorWhite = "#FFFFFF"
)

var (
	Base = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorWhite))

	Key = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorWhite))

	Dim = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorWhite))

	Heading = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorWhite)).
		Bold(true)

	Subtle = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorWhite))

	Accent = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorWhite))

	Portrait = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorWhite))

	Page = lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorWhite)).
		Padding(2, 4)
)
