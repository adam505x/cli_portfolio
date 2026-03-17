package web

import (
	"embed"
	"encoding/json"
	"io"
	"io/fs"
	"net/http"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
	"portfolio/tui"
)

//go:embed static
var staticFiles embed.FS

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func NewServer(addr string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", handleWS)

	sub, _ := fs.Sub(staticFiles, "static")
	mux.Handle("/", http.FileServer(http.FS(sub)))

	return &http.Server{Addr: addr, Handler: mux}
}

// wsReader bridges incoming WebSocket messages to an io.Reader for bubbletea.
type wsReader struct {
	msgs chan []byte
	mu   sync.Mutex
	buf  []byte
}

func (r *wsReader) Read(p []byte) (n int, err error) {
	r.mu.Lock()
	if len(r.buf) > 0 {
		n = copy(p, r.buf)
		r.buf = r.buf[n:]
		r.mu.Unlock()
		return n, nil
	}
	r.mu.Unlock()

	msg, ok := <-r.msgs
	if !ok {
		return 0, io.EOF
	}

	r.mu.Lock()
	n = copy(p, msg)
	if n < len(msg) {
		r.buf = append(r.buf, msg[n:]...)
	}
	r.mu.Unlock()
	return n, nil
}

// wsWriter bridges bubbletea's rendered output to outgoing WebSocket messages.
type wsWriter struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

func (w *wsWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if err = w.conn.WriteMessage(websocket.BinaryMessage, p); err != nil {
		return 0, err
	}
	return len(p), nil
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("websocket upgrade failed", "err", err)
		return
	}
	defer conn.Close()
	log.Info("web client connected", "addr", r.RemoteAddr)

	reader := &wsReader{msgs: make(chan []byte, 64)}
	writer := &wsWriter{conn: conn}

	m := tui.NewModel(220, 50)
	p := tea.NewProgram(m,
		tea.WithInput(reader),
		tea.WithOutput(writer),
		tea.WithAltScreen(),
	)

	// Read WebSocket messages: resize events are dispatched as tea messages,
	// everything else (keystrokes) is forwarded to bubbletea's input reader.
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				close(reader.msgs)
				p.Kill()
				return
			}

			var ev struct {
				Type string `json:"type"`
				Cols int    `json:"cols"`
				Rows int    `json:"rows"`
			}
			if json.Unmarshal(msg, &ev) == nil && ev.Type == "resize" {
				p.Send(tea.WindowSizeMsg{Width: ev.Cols, Height: ev.Rows})
				continue
			}

			select {
			case reader.msgs <- msg:
			default:
			}
		}
	}()

	if _, err := p.Run(); err != nil {
		log.Error("bubbletea error", "err", err)
	}
	log.Info("web client disconnected", "addr", r.RemoteAddr)
}
