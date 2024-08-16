package usecase

import (
	"chatroom/internal/domain"
	"chatroom/internal/repository/wsrepo"
	"log"
)

type WebSocketUsecase struct {
	repo wsrepo.ConnectionRepository
}

func NewWebSocketUsecase(repo wsrepo.ConnectionRepository) *WebSocketUsecase {
	return &WebSocketUsecase{repo: repo}
}

func (uc *WebSocketUsecase) HandleNewConnection(conn *domain.Connection) {
	uc.repo.AddConnection(conn)
}

func (uc *WebSocketUsecase) HandleDisconnect(conn *domain.Connection) {
	uc.repo.RemoveConnection(conn)
}

func (uc *WebSocketUsecase) BroadcastMessage(sender *domain.Connection, msg *domain.Message) {
	connections := uc.repo.GetConnections()
	for _, c := range connections {
		if c != sender {
			log.Println("teste aquii", c, string(msg.Content))
			err := c.Conn.WriteMessage(1, msg.Content)
			if err != nil {
				log.Println("Error broadcasting message:", err)
			}
		}
	}
}
