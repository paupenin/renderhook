package api

import (
	"fmt"
	"net/http"

	"github.com/paupenin/web2image/backend/browser"
)

type Server struct {
	Config      *ServerConfig
	browserPool *browser.BrowserPool
}

// Creates a new server
func NewServer(config *ServerConfig) *Server {
	return &Server{
		Config:      config,
		browserPool: browser.NewBrowserPool(3),
	}
}

// Starts the server
func (s *Server) Start() {
	fmt.Println("Starting server")

	// Initialize the browser pool
	s.browserPool.Init()

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
