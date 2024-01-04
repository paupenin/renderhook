package main

import "github.com/paupenin/web2image/backend/api"

func main() {
	// Build and start the server
	api.NewServer(
		api.NewServerConfig(),
	).Start()
}
