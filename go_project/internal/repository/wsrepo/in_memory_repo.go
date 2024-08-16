package wsrepo

import (
	"chatroom/internal/domain"
	"log"
	"sync"
)

type ConnectionRepository interface {
	AddConnection(conn *domain.Connection)
	RemoveConnection(conn *domain.Connection)
	GetConnections() []*domain.Connection
}

type InMemoryConnectionRepo struct {
	connections map[*domain.Connection]bool
	mu          sync.Mutex
}

func NewInMemoryConnectionRepo() *InMemoryConnectionRepo {
	return &InMemoryConnectionRepo{
		connections: make(map[*domain.Connection]bool),
	}
}

func (repo *InMemoryConnectionRepo) AddConnection(conn *domain.Connection) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.connections[conn] = true
	log.Println("connections", repo.connections)
}

func (repo *InMemoryConnectionRepo) RemoveConnection(conn *domain.Connection) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	delete(repo.connections, conn)
	log.Println("connections", repo.connections)
}

func (repo *InMemoryConnectionRepo) GetConnections() []*domain.Connection {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	conns := make([]*domain.Connection, 0, len(repo.connections))
	for conn := range repo.connections {
		conns = append(conns, conn)
	}
	return conns
}
