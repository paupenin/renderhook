package store

import (
	"os"

	"github.com/paupenin/web2image/backend/config"
)

type FileStoreFS struct {
	config *config.FileStoreFSConfig
}

func NewFileStoreFS(config *config.FileStoreFSConfig) *FileStoreFS {
	return &FileStoreFS{
		config: config,
	}
}

// ShouldServeStatic returns whether the file store should serve static files
func (s *FileStoreFS) ShouldServeStatic() bool {
	return true
}

// GetStaticPath gets the path to the static file store
func (s *FileStoreFS) GetStaticPath() string {
	return s.config.Directory
}

// GetURL gets the public URL to the file store
func (s *FileStoreFS) GetURL() string {
	return s.config.PublicURL
}

// StoreFile stores a file
func (s *FileStoreFS) StoreFile(filename string, image []byte) error {
	// Create directory if it doesn't exist
	err := os.MkdirAll(s.config.Directory, 0777)

	if err != nil {
		return err
	}

	// Write file
	err = os.WriteFile(s.config.Directory+"/"+filename, image, 0666)

	if err != nil {
		return err
	}

	return nil
}

// GetFileURL gets the URL of an image
func (s *FileStoreFS) GetFileURL(filename string) string {
	return s.config.PublicURL + "/" + filename
}

// DeleteImage deletes an image
func (s *FileStoreFS) DeleteFile(filename string) error {
	// Delete file
	err := os.Remove(s.config.Directory + "/" + filename)

	if err != nil {
		return err
	}

	return nil
}
