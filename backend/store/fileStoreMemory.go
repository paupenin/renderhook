package store

import (
	"fmt"
	"sync"
)

type FileStoreMemory struct {
	files map[string][]byte
	mu    sync.RWMutex // Protects the files map
}

func NewFileStoreMemory() *FileStoreMemory {
	return &FileStoreMemory{
		files: make(map[string][]byte),
	}
}

func (s *FileStoreMemory) ShouldServeStatic() bool {
	return false
}

func (s *FileStoreMemory) GetStaticPath() string {
	return "" // Not applicable for in-memory store
}

func (s *FileStoreMemory) StoreFile(filename string, file []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.files[filename]; exists {
		return fmt.Errorf("file already exists: %s", filename)
	}

	s.files[filename] = file
	return nil
}

func (s *FileStoreMemory) GetFileURL(filename string) string {
	return "memory://" + filename
}

func (s *FileStoreMemory) DeleteFile(filename string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.files[filename]; !exists {
		return fmt.Errorf("file does not exist: %s", filename)
	}

	delete(s.files, filename)
	return nil
}
