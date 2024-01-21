package store

import (
	"testing"
)

func TestFileStoreMemory(t *testing.T) {
	fs := NewFileStoreMemory()

	// Test ShouldServeStatic
	if fs.ShouldServeStatic() {
		t.Error("ShouldServeStatic should return false")
	}

	// Test GetStaticPath
	if fs.GetStaticPath() != "" {
		t.Error("GetStaticPath should return an empty string")
	}

	// Test data
	filename := "test.txt"
	content := []byte("hello world")

	// Test StoreFile
	if err := fs.StoreFile(filename, content); err != nil {
		t.Errorf("StoreFile failed: %v", err)
	}

	// Test storing a file with a duplicate name
	if err := fs.StoreFile(filename, content); err == nil {
		t.Error("StoreFile should fail when storing a file with a duplicate name")
	}

	// Test GetFileURL
	expectedURL := "memory://" + filename
	if url := fs.GetFileURL(filename); url != expectedURL {
		t.Errorf("GetFileURL returned wrong URL: got %v, want %v", url, expectedURL)
	}

	// Test DeleteFile
	if err := fs.DeleteFile(filename); err != nil {
		t.Errorf("DeleteFile failed: %v", err)
	}

	// Test deleting a non-existent file
	if err := fs.DeleteFile(filename); err == nil {
		t.Error("DeleteFile should fail when deleting a non-existent file")
	}
}
