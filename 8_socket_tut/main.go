package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("New Connection From Client: ", ws.RemoteAddr())

	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buff := make([]byte, 1024)
	for {
		n, err := ws.Read(buff)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Read error: ", err)
			continue
		}
		msg := buff[:n]
		s.broadcast(msg)
	}
}
func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/feed", websocket.Handler(server.handleFeed))
	http.ListenAndServe(":3000", nil)
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("Write error: ", err)
			}
		}(ws)
	}
}

func (s *Server) handleFeed(ws *websocket.Conn) {
	fmt.Println("New Client From ", ws.RemoteAddr(), " is accessing live feed")
	for {
		payload := fmt.Sprintf("Current Runrate -> %d ", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(3 * time.Second)
	}
}
