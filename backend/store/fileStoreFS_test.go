package store

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/paupenin/renderhook/backend/config"
)

func setupTestFileStoreFS(t *testing.T) (*FileStoreFS, string) {
	t.Helper()

	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "testfilestorefs")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}

	// Clean up after the test
	defer os.RemoveAll(tempDir)

	// Setup configuration
	cfg := &config.FileStoreFSConfig{
		Directory: tempDir,
		PublicURL: "http://fs.test",
	}

	return NewFileStoreFS(cfg), tempDir
}

func TestFileStoreFS_ShouldServeStatic(t *testing.T) {
	fs, _ := setupTestFileStoreFS(t)

	if !fs.ShouldServeStatic() {
		t.Error("ShouldServeStatic should return true")
	}
}

func TestFileStoreFS_GetStaticPath(t *testing.T) {
	fs, tempDir := setupTestFileStoreFS(t)

	if fs.GetStaticPath() != tempDir {
		t.Errorf("GetStaticPath returned wrong directory: got %v, want %v", fs.GetStaticPath(), tempDir)
	}
}

func TestFileStoreFS_StoreFile(t *testing.T) {
	fs, tempDir := setupTestFileStoreFS(t)
	defer os.RemoveAll(tempDir) // Clean up after the test

	filename := "test.txt"
	content := []byte("hello world")

	if err := fs.StoreFile(filename, content); err != nil {
		t.Errorf("StoreFile failed: %v", err)
	}

	// Verify file contents
	filePath := filepath.Join(tempDir, filename)
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Failed to read stored file: %v", err)
	}
	if string(fileContent) != string(content) {
		t.Errorf("Stored file content mismatch: got %v, want %v", string(fileContent), string(content))
	}
}

func TestFileStoreFS_GetFileURL(t *testing.T) {
	fs, _ := setupTestFileStoreFS(t)

	filename := "test.txt"
	expectedURL := "http://fs.test/" + filename

	if url := fs.GetFileURL(filename); url != expectedURL {
		t.Errorf("GetFileURL returned wrong URL: got %v, want %v", url, expectedURL)
	}
}

func TestFileStoreFS_DeleteFile(t *testing.T) {
	fs, tempDir := setupTestFileStoreFS(t)
	defer os.RemoveAll(tempDir)

	filename := "test.txt"
	content := []byte("hello world")

	// Store a file first
	if err := fs.StoreFile(filename, content); err != nil {
		t.Fatalf("StoreFile failed: %v", err)
	}

	// Then try to delete it
	if err := fs.DeleteFile(filename); err != nil {
		t.Errorf("DeleteFile failed: %v", err)
	}

	// Verify file is deleted
	if _, err := os.Stat(filepath.Join(tempDir, filename)); !os.IsNotExist(err) {
		t.Errorf("File should be deleted, but still exists")
	}
}
