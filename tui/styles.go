package tui

import (
	_ "embed"

	"github.com/charmbracelet/lipgloss"
)

const (
	colorBg    = "#000000"
	colorWhite = "#FFFFFF"
)

//go:embed assets/name.txt
var nameArtRaw string

var (
	baseStyle = lipgloss.NewStyle().
			Background(lipgloss.Color(colorBg)).
			Foreground(lipgloss.Color(colorWhite))

	keyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorWhite)).
			Background(lipgloss.Color(colorBg))

	dimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorWhite)).
			Background(lipgloss.Color(colorBg))

	headingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorWhite)).
			Background(lipgloss.Color(colorBg)).
			Bold(true)

	subtleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorWhite)).
			Background(lipgloss.Color(colorBg))

	accentStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorWhite)).
			Background(lipgloss.Color(colorBg))

	nameArt = baseStyle.Render(nameArtRaw)

	portraitStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorWhite)).
			Background(lipgloss.Color(colorBg))

	pageStyle = lipgloss.NewStyle().
			Background(lipgloss.Color(colorBg)).
			Padding(2, 4)
)
