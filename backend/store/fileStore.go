package store

import (
	"github.com/paupenin/renderhook/backend/config"
)

// Interface for all file stores
type FileStore interface {
	// ShouldServeStatic returns whether the file store should serve static files
	ShouldServeStatic() bool
	// GetStaticPath gets the path to the static file store directory (if applicable)
	GetStaticPath() string
	// StoreFile stores a file
	StoreFile(filename string, file []byte) error
	// GetFileURL gets the public URL of a file
	GetFileURL(filename string) string
	// DeleteFile deletes an file
	DeleteFile(filename string) error
}

func NewFileStore(c config.FileStoreConfig) FileStore {
	switch c := c.(type) {
	// File Store S3 (production)
	case *config.FileStoreS3Config:
		return NewFileStoreS3(c)

	// File Store FS (development)
	case *config.FileStoreFSConfig:
		return NewFileStoreFS(c)

	// File Store Memory (testing)
	default:
		return NewFileStoreMemory()
	}
}
