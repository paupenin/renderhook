package main

import (
	"github.com/paupenin/renderhook/backend/api"
	"github.com/paupenin/renderhook/backend/config"
)

func main() {
	// Build and start the server
	api.NewServer(
		config.NewServerConfig(),
	).Start()
}
