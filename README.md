# cli-portfolio

A terminal portfolio served over SSH, built with the [Charm.sh](https://charm.sh) ecosystem.

## Stack

- [bubbletea](https://github.com/charmbracelet/bubbletea) — TUI framework
- [lipgloss](https://github.com/charmbracelet/lipgloss) — terminal styling
- [wish](https://github.com/charmbracelet/wish) — SSH server

## Setup

### 1. Install Go

Download from https://go.dev/dl/ (1.22+).

### 2. Install dependencies

```bash
cd cli_portfolio
go mod tidy
```

### 3. Run

```bash
go run main.go
```

The server starts on port **2222** and generates `.ssh/id_ed25519` automatically on first run.

### 4. Connect

```bash
ssh localhost -p 2222
```

Or from another machine:

```bash
ssh yourserver.com -p 2222
```

## Navigation

| Key   | Action        |
|-------|---------------|
| `1`   | Projects      |
| `2`   | Blog posts    |
| `3`   | About         |
| `esc` | Back to home  |
| `q`   | Quit          |

## Customisation

| File                  | What to change                          |
|-----------------------|-----------------------------------------|
| `tui/app.go`          | Projects, blog posts, about info        |
| `tui/styles.go`       | Colours, name ASCII art (`nameArt` var) |
| `tui/portrait.go`     | ASCII portrait                          |
| `main.go`             | Port, host key path                     |

### Changing the name art

The `nameArt` variable in `tui/styles.go` is a plain string constant.
Generate a new one with [figlet](http://www.figlet.org/):

```bash
figlet -f banner "Your Name"
```

Then paste it into the backtick string.

## Deployment

Any Linux box works. Run it as a systemd service or inside a tmux session.
Point port 2222 (or 22) to the process in your firewall/reverse proxy.
The `.ssh/` directory holding the host key should persist across restarts.
