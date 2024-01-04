package browser

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/ysmood/gson"
)

// Browser is the render engine
type Browser struct {
	rod *rod.Browser
}

// NewBrowser creates a new browser
func NewBrowser() *Browser {
	return &Browser{}
}

// Connect connects to the browser
func (b *Browser) Connect() error {
	browser := rod.New().MustConnect()

	b.rod = browser

	return nil
}

// Close closes the browser
func (b *Browser) Close() error {
	b.rod.MustClose()

	return nil
}

// RenderHTML renders HTML
func (b *Browser) RenderHTML(html string) ([]byte, error) {
	fmt.Println("Rendering HTML string")

	// Load page
	page := b.rod.MustPage()

	// Set viewport
	page.MustSetViewport(500, 500, 2, false)

	// Set HTML content
	page.SetDocumentContent(html)

	// Wait for page to load
	page.WaitLoad()

	// Return screenshot
	return page.Screenshot(true, &proto.PageCaptureScreenshot{
		Format:  proto.PageCaptureScreenshotFormatJpeg,
		Quality: gson.Int(99),
	})
}

// RenderURL renders a URL
func (b *Browser) RenderURL(url string) ([]byte, error) {
	fmt.Println("Rendering URL: ", url)

	// Load page
	page := b.rod.MustPage(url)

	// Set viewport
	page.SetViewport(&proto.EmulationSetDeviceMetricsOverride{
		Width:             320,
		DeviceScaleFactor: 2,
		Mobile:            false,
	})

	// Wait for page to load
	page.WaitLoad()

	// Return screenshot
	return page.Screenshot(true, &proto.PageCaptureScreenshot{
		Format:  proto.PageCaptureScreenshotFormatJpeg,
		Quality: gson.Int(99),
	})
}
