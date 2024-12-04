package echo

import "sync"

/*
服务端：维护一个所有活动中的clients，接收注册和注销两个信号，并且能够广播消息给客户端
*/
//var MyServer = NewServer()

type Server struct {
	mutex      *sync.Mutex
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewServer() *Server {
	return &Server{
		mutex:      &sync.Mutex{},
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (s *Server) Run() {
	for {
		select {
		case client := <-s.register:
			s.mutex.Lock()
			s.clients[client] = true
			s.mutex.Unlock()
		case client := <-s.unregister:
			if _, ok := s.clients[client]; ok {
				s.mutex.Lock()
				delete(s.clients, client)
				s.mutex.Unlock()
				close(client.Send)
			}
		case message := <-s.broadcast:
			for client := range s.clients {
				select {
				case client.Send <- message:
				}
			}
		}
	}
}

// 群组通信
func (s *Server) Broadcast(message []byte) {
	s.broadcast <- message
}

func (s *Server) Register(client *Client) {
	s.register <- client
}

// 消息持久化
func (s *Server) SaveMessage(message []byte) {
	//TODO 将消息持久化到数据库

}
