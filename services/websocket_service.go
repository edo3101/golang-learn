package services

import (
	"learn/models"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type WebsocketService struct {
	Clients   map[*websocket.Conn]bool
	Broadcast chan models.WebsocketMessage
	Upgrader  websocket.Upgrader
	mu        sync.Mutex
}

func NewWebsocketService() *WebsocketService {
	return &WebsocketService{
		Clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan models.WebsocketMessage),
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (s *WebsocketService) AddClient(conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Clients[conn] = true
}

func (s *WebsocketService) RemoveClient(conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Clients[conn]; ok {
		delete(s.Clients, conn)
		conn.Close()
	}
}

func (s *WebsocketService) BroadcastService(message models.WebsocketMessage) {
	s.Broadcast <- message
}

func (s *WebsocketService) HandleMessage() {
	for {
		message := <-s.Broadcast
		s.mu.Lock()
		for client := range s.Clients {
			err := client.WriteJSON(message)
			if err != nil {
				client.Close()
				delete(s.Clients, client)
			}
		}
		s.mu.Unlock()
	}
}
