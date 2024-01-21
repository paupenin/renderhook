package store

import "github.com/paupenin/web2image/backend/config"

type FileStoreS3 struct {
	config *config.FileStoreS3Config
}

func NewFileStoreS3(c *config.FileStoreS3Config) *FileStoreS3 {
	return &FileStoreS3{
		config: c,
	}
}

// ShouldServeStatic returns whether the file store should serve static files
func (s *FileStoreS3) ShouldServeStatic() bool {
	return false
}

// GetStaticPath gets the path to the static file store
func (s *FileStoreS3) GetStaticPath() string {
	return ""
}

// GetURL gets the public URL to the file store
func (s *FileStoreS3) GetURL() string {
	return ""
}

// StoreFile stores a file
func (s *FileStoreS3) StoreFile(filename string, file []byte) error {
	return nil
}

// GetFileURL gets the URL of an image
func (s *FileStoreS3) GetFileURL(filename string) string {
	return ""
}

// DeleteImage deletes an image
func (s *FileStoreS3) DeleteFile(filename string) error {
	return nil
}
