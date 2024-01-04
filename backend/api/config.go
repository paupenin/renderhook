package api

type ServerConfig struct {
	// The server's host
	Host string
	// The server's port
	Port int
	// Whether to use SSL
	SSL bool
}

// Creates a new server config
func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Host: "localhost",
		Port: 8080,
		SSL:  false,
	}
}

// Sets the server's port
func (c *ServerConfig) SetPort(port int) *ServerConfig {
	c.Port = port
	return c
}

// Sets the server's host
func (c *ServerConfig) SetHost(host string) *ServerConfig {
	c.Host = host
	return c
}

// Sets whether to use SSL
func (c *ServerConfig) SetSSL(ssl bool) *ServerConfig {
	c.SSL = ssl
	return c
}
