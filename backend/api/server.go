package api

import (
	"fmt"
	"net/http"

	"github.com/paupenin/renderhook/backend/browser"
	"github.com/paupenin/renderhook/backend/config"
	"github.com/paupenin/renderhook/backend/store"
)

type Server struct {
	config      *config.ServerConfig
	browserPool *browser.BrowserPool
	imageStore  store.FileStore
}

// Creates a new server
func NewServer(c config.ServerConfig) *Server {
	return &Server{
		config:      &c,
		browserPool: browser.NewBrowserPool(c.BrowserPool),
		imageStore:  store.NewFileStore(c.Storage),
	}
}

// Creates a new testServer
func NewTestServer() *Server {
	return &Server{
		config: &config.ServerConfig{},
		browserPool: browser.NewBrowserPool(
			config.BrowserPoolConfig{
				MaxBrowsers:     1,
				MaxBrowserPages: 1,
			},
		),
		imageStore: store.NewFileStore(nil), // Memory store
	}
}

// Starts the server
func (s *Server) Start() {
	fmt.Println("Starting server")

	// Initialize the browser pool
	s.browserPool.Init()

	fmt.Println("Server listening on " + s.config.GetURL())

	http.ListenAndServe(s.config.GetAddress(), s.initRouter())
}

// Stops the server
func (s *Server) Stop() {
	fmt.Println("Stopping server")
}
