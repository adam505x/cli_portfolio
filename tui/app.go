package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"portfolio/tui/pages"
	"portfolio/tui/router"
	"portfolio/tui/services"
)

// ── Async messages ────────────────────────────────────────────────────────────

type lastFMMsg struct{ tracks []services.Track }
type lastFMErrMsg struct{}
type letterboxdMsg struct{ films []services.Film }
type tickMsg struct{}

// ── Model ─────────────────────────────────────────────────────────────────────

type Model struct {
	stack     []router.Page
	width     int
	height    int
	animFrame int
}

func NewModel(w, h int) Model {
	pages.StarsArt = StarsArt
	pages.GalaxyArt = GalaxyArt
	pages.CatArt = CatArt
	pages.HeadphonesArt = HeadphonesArt
	return Model{
		stack:  []router.Page{pages.NewHome(Portrait, NameArt)},
		width:  w,
		height: h,
	}
}

// Init kicks off background data fetches and the animation ticker.
func (m Model) Init() tea.Cmd {
	return tea.Batch(fetchLastFM(), parseLetterboxd(), tick())
}

// ── Startup commands ──────────────────────────────────────────────────────────

func tick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func fetchLastFM() tea.Cmd {
	return func() tea.Msg {
		tracks, err := services.FetchTopTracks()
		if err != nil {
			return lastFMErrMsg{}
		}
		return lastFMMsg{tracks: tracks}
	}
}

func parseLetterboxd() tea.Cmd {
	return func() tea.Msg {
		films, err := services.ParseDiary(DiaryCSV, 5)
		if err != nil {
			return letterboxdMsg{}
		}
		return letterboxdMsg{films: films}
	}
}

// ── Update ────────────────────────────────────────────────────────────────────

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		pages.TermHeight = msg.Height
		pages.TermWidth = msg.Width

	case lastFMMsg:
		pages.TopTracks = msg.tracks

	case lastFMErrMsg:
		pages.TopTracks = []services.Track{{Name: "—", Artist: "", PlayCount: ""}}

	case letterboxdMsg:
		pages.RecentFilms = msg.films

	case tickMsg:
		m.animFrame++
		pages.AnimFrame = m.animFrame
		return m, tick()

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		current := m.stack[len(m.stack)-1]
		switch a := current.Update(msg.String()).(type) {
		case router.Stay:
			m.stack[len(m.stack)-1] = a.Page
		case router.Push:
			m.stack = append(m.stack, a.Page)
		case router.Pop:
			if len(m.stack) > 1 {
				m.stack = m.stack[:len(m.stack)-1]
			}
		case router.Quit:
			return m, tea.Quit
		}
	}

	return m, nil
}

// ── View ──────────────────────────────────────────────────────────────────────

func (m Model) View() string {
	return m.stack[len(m.stack)-1].View()
}
