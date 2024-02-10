package store

import (
	"testing"

	"github.com/paupenin/renderhook/backend/config"
)

func setupTestFileStoreS3(t *testing.T) (*FileStoreS3, *MockS3Client) {
	t.Helper()

	// Setup configuration
	cfg := &config.FileStoreS3Config{
		BucketName: "testbucket",
		Region:     "testregion",
		Endpoint:   "testendpoint",
		AccessKey:  "testaccesskey",
		SecretKey:  "testsecretkey",
	}

	// Create a mock S3 client
	client := NewMockS3Client()

	return &FileStoreS3{
		config: cfg,
		client: client,
	}, client
}

func TestFileStoreS3_ShouldServeStatic(t *testing.T) {
	fs, _ := setupTestFileStoreS3(t)

	if fs.ShouldServeStatic() {
		t.Error("ShouldServeStatic should return false")
	}
}

func TestFileStoreS3_GetStaticPath(t *testing.T) {
	fs, _ := setupTestFileStoreS3(t)

	if fs.GetStaticPath() != "" {
		t.Errorf("GetStaticPath should return an empty string")
	}
}

func TestFileStoreS3_StoreFile(t *testing.T) {
	fs, client := setupTestFileStoreS3(t)

	filename := "test.txt"
	content := []byte("hello world")

	if err := fs.StoreFile(filename, content); err != nil {
		t.Errorf("StoreFile failed: %v", err)
	}

	// Verify file contents
	if _, exists := client.StoredObjects[filename]; !exists {
		t.Errorf("Stored file not found")
	}
	if string(client.StoredObjects[filename]) != string(content) {
		t.Errorf("Stored file content mismatch: got %v, want %v", string(client.StoredObjects[filename]), string(content))
	}
}

func TestFileStoreS3_GetFileURL(t *testing.T) {
	fs, _ := setupTestFileStoreS3(t)

	filename := "test.txt"
	expectedURL := "testendpoint/test.txt"

	if url := fs.GetFileURL(filename); url != expectedURL {
		t.Errorf("GetFileURL returned wrong URL: got %v, want %v", url, expectedURL)
	}
}

func TestFileStoreS3_DeleteFile(t *testing.T) {
	fs, client := setupTestFileStoreS3(t)

	filename := "test.txt"
	content := []byte("hello world")

	// Store a file
	client.StoredObjects[filename] = content

	if err := fs.DeleteFile(filename); err != nil {
		t.Errorf("DeleteFile failed: %v", err)
	}

	// Verify file is deleted
	if _, exists := client.StoredObjects[filename]; exists {
		t.Errorf("Stored file not deleted")
	}
}
