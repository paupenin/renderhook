package browser

import "github.com/paupenin/renderhook/backend/config"

// Browser pool
type BrowserPool struct {
	// Amount of desired browsers
	amount int
	// Browsers
	browsers []*Browser
}

// NewBrowserPool creates a new browser pool
func NewBrowserPool(c config.BrowserPoolConfig) *BrowserPool {
	return &BrowserPool{
		amount: c.MaxBrowsers,
	}
}

// Init initializes the browser pool
func (bp *BrowserPool) Init() error {
	for i := 0; i < bp.amount; i++ {
		browser := NewBrowser()

		err := browser.Init()

		if err != nil {
			return err
		}

		bp.browsers = append(bp.browsers, browser)
	}

	return nil
}

// Destroy destroys the browser pool
func (bp *BrowserPool) Destroy() error {
	for _, browser := range bp.browsers {
		browser.Destroy()
	}

	return nil
}

// GetAmount gets the amount of desired browsers in the pool
func (bp *BrowserPool) GetAmount() int {
	return bp.amount
}

// GetAvailableAmount gets the amount of available browsers in the pool
func (bp *BrowserPool) GetAvailableAmount() int {
	return len(bp.browsers)
}

func (bp *BrowserPool) IsReady() bool {
	return bp.GetAvailableAmount() > 0
}

// GetBrowser gets a browser from the pool
func (bp *BrowserPool) GetBrowser() *Browser {
	return bp.browsers[0]
}

// We don't put browsers back into the pool, we get a new one each time (?)
