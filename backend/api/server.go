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

	fmt.Println("Server listening on " + s.GetURL())

	http.ListenAndServe(s.GetAddress(), s.initRouter())
}

// Stops the server
func (s *Server) Stop() {
	fmt.Println("Stopping server")
}

// Get Address
func (s *Server) GetAddress() string {
	return s.Config.Host + ":" + fmt.Sprint(s.Config.Port)
}

// Get URL
func (s *Server) GetURL() string {
	if s.Config.SSL {
		return "https://" + s.GetAddress()
	}

	return "http://" + s.GetAddress()
}
