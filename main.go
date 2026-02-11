package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Message struct {
	from    string // sender address
	payload []byte // message content
	at 		time.Time // timestamp
}
type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     chan struct{}
	msgch      chan Message // buffered to avoid blocking
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		msgch:      make(chan Message, 10),
	}
}
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.ln = ln
	go s.acceptLoop()
	<-s.quitch // wait for shutdown signal
	close(s.msgch)
	return nil
}
func (s *Server) acceptLoop() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("new connection to the server:", conn.RemoteAddr())
		go s.readLoop(conn) // handle each client in its own goroutine
	}
}
func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		// send msg to channel and confirm receipt
		s.msgch <- Message{
			from:    conn.RemoteAddr().String(),
			payload: buf[:n],
			at: time.Now(),
		}
		_, _ = conn.Write(buf[:n])
	}
}
func main() {
	server := NewServer(":3000")
	// handle incoming messages
	go func() {
		for msg := range server.msgch {
			fmt.Printf("[%s] %s: %s\n",
				msg.at.Format(time.RFC3339),
				msg.from,
				string(msg.payload),
			)
		}
	}()
	log.Fatal(server.Start())
}
