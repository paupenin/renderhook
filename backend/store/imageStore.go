package store

import (
	"os"
)

// ImageStore is the interface for the image store
type ImageStore interface {
	// StoreImage stores an image
	StoreImage(filename string, image []byte) error
	// GetImage gets an image
	GetImage(filename string) ([]byte, error)
	// GetImagePath gets the path of an image
	GetImagePath(filename string) string
	// DeleteImage deletes an image
	DeleteImage(filename string) error
}

// ImageStoreConfig is the configuration for the image store
type ImageStoreConfig struct {
	// Path is the path to the image store
	Path string
}

// ImageStore is the image store
type imageStore struct {
	Config ImageStoreConfig
}

// NewImageStore creates a new image store
func NewImageStore(config ImageStoreConfig) ImageStore {
	return &imageStore{
		Config: config,
	}
}

// StoreImage stores an image
func (s *imageStore) StoreImage(filename string, image []byte) error {
	// Create directory if it doesn't exist
	err := os.MkdirAll(s.Config.Path, 0777)

	if err != nil {
		return err
	}

	// Write file
	err = os.WriteFile(s.Config.Path+"/"+filename, image, 0666)

	if err != nil {
		return err
	}

	return nil
}

// GetImage gets an image
func (s *imageStore) GetImage(filename string) ([]byte, error) {
	// Read file
	image, err := os.ReadFile(s.Config.Path + "/" + filename)

	if err != nil {
		return nil, err
	}

	return image, nil
}

// GetImageURL gets the URL of an image
func (s *imageStore) GetImagePath(filename string) string {
	return s.Config.Path + "/" + filename
}

// DeleteImage deletes an image
func (s *imageStore) DeleteImage(filename string) error {
	// Delete file
	err := os.Remove(s.Config.Path + "/" + filename)

	if err != nil {
		return err
	}

	return nil
}
