package store

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/paupenin/renderhook/backend/config"
)

type FileStoreFS struct {
	config *config.FileStoreFSConfig
}

func NewFileStoreFS(cfg *config.FileStoreFSConfig) *FileStoreFS {
	return &FileStoreFS{
		config: cfg,
	}
}

func (s *FileStoreFS) ShouldServeStatic() bool {
	return true
}

func (s *FileStoreFS) GetStaticPath() string {
	return s.config.Directory
}

func (s *FileStoreFS) StoreFile(filename string, image []byte) error {
	filePath := filepath.Join(s.config.Directory, filename)

	// Create directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	// Write file
	if err := os.WriteFile(filePath, image, 0644); err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

func (s *FileStoreFS) GetFileURL(filename string) string {
	return s.config.PublicURL + "/" + filename
}

func (s *FileStoreFS) DeleteFile(filename string) error {
	filePath := filepath.Join(s.config.Directory, filename)

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("error deleting file: %w", err)
	}

	return nil
}
