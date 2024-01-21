package store

import (
	"log"

	"github.com/paupenin/web2image/backend/config"
)

// Interface for all file stores
type FileStore interface {
	// ShouldServeStatic returns whether the file store should serve static files
	ShouldServeStatic() bool
	// GetStaticPath gets the path to the static file store
	GetStaticPath() string
	// GetStaticURL gets the public URL to the file store
	GetURL() string
	// StoreFile stores a file
	StoreFile(filename string, file []byte) error
	// GetFileURL gets the public URL of a file
	GetFileURL(filename string) string
	// DeleteFile deletes an file
	DeleteFile(filename string) error
}

func NewFileStore(c config.FileStoreConfig) FileStore {
	switch c := c.(type) {
	case *config.FileStoreFSConfig:
		return NewFileStoreFS(c)

	case *config.FileStoreS3Config:
		return NewFileStoreS3(c)
	}

	log.Fatal("Unknown file store type")
	return nil
}
