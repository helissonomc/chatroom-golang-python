package wsdelivery

import (
	"chatroom/internal/domain"
	"chatroom/internal/usecase"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WebSocketHandler struct {
	usecase *usecase.WebSocketUsecase
}

func NewWebSocketHandler(usecase *usecase.WebSocketUsecase) *WebSocketHandler {
	return &WebSocketHandler{usecase: usecase}
}

func (h *WebSocketHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	connection := &domain.Connection{Conn: conn}
	h.usecase.HandleNewConnection(connection)
	defer func() {
		log.Println("defer executed")
		h.usecase.HandleDisconnect(connection)
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		log.Println("aaa", string(msg), err)
		if err != nil {
			log.Println(err)
			break
		}

		message := &domain.Message{Content: msg}
		h.usecase.BroadcastMessage(connection, message)
	}
}
