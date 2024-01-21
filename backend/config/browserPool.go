package config

type BrowserPoolConfig struct {
	// MaxBrowsers is the maximum number of browsers to keep open
	MaxBrowsers int
	// MaxBrowserPages is the maximum number of pages to keep open per browser
	MaxBrowserPages int
}

func NewBrowserPoolConfig() BrowserPoolConfig {
	return BrowserPoolConfig{
		MaxBrowsers:     EnvInt("BROWSER_POOL_MAX_BROWSERS", 3),
		MaxBrowserPages: EnvInt("BROWSER_POOL_MAX_BROWSER_PAGES", 3),
	}
}
