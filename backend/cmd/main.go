package main

import (
	"github.com/paupenin/web2image/backend/api"
	"github.com/paupenin/web2image/backend/config"
)

func main() {
	// Build and start the server
	api.NewServer(
		config.NewServerConfig(),
	).Start()
}
