package api

import (
	"fmt"
	"net/http"
)

type Server struct {
	Config *ServerConfig
}

// Creates a new server
func NewServer(config *ServerConfig) *Server {
	return &Server{
		Config: config,
	}
}

// Starts the server
func (s *Server) Start() {
	fmt.Println("Starting server")

	addr := s.Config.Host + ":" + fmt.Sprint(s.Config.Port)

	fmt.Println("Server listening on http://" + addr)

	http.ListenAndServe(addr, s.initRouter())
}

// Stops the server
func (s *Server) Stop() {
	fmt.Println("Stopping server")
}
