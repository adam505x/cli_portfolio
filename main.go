package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
	"github.com/joho/godotenv"
	"portfolio/tui"
	"portfolio/web"
)

const (
	host    = "0.0.0.0"
	sshPort = 2222
	webPort = 8080
	keyPath = ".ssh/id_ed25519"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Warn("no .env file found, falling back to environment variables")
	}

	// Local mode: run TUI directly in the current terminal (no SSH)
	if len(os.Args) > 1 && os.Args[1] == "-local" {
		m := tui.NewModel(220, 50)
		p := tea.NewProgram(m, tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			log.Error("local run error", "error", err)
			os.Exit(1)
		}
		return
	}

	// SSH server
	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, strconv.Itoa(sshPort))),
		wish.WithHostKeyPath(keyPath),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			activeterm.Middleware(),
			lm.Middleware(),
		),
	)
	if err != nil {
		log.Error("Could not create SSH server", "error", err)
		os.Exit(1)
	}

	// Web server
	webServer := web.NewServer(":" + strconv.Itoa(webPort))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Info("SSH server starting", "host", host, "port", sshPort)
	log.Info("Web server starting", "host", host, "port", webPort)

	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("SSH server error", "error", err)
			done <- nil
		}
	}()

	go func() {
		if err := webServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("Web server error", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("SSH shutdown error", "error", err)
	}
	if err := webServer.Shutdown(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error("Web shutdown error", "error", err)
	}
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	w, h := pty.Window.Width, pty.Window.Height
	if w == 0 {
		w = 220
	}
	if h == 0 {
		h = 50
	}
	m := tui.NewModel(w, h)
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}
