package store

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/service/s3"
)

// MockS3Client is a mock of S3API interface
type MockS3Client struct {
	StoredObjects map[string][]byte
}

func NewMockS3Client() *MockS3Client {
	return &MockS3Client{
		StoredObjects: make(map[string][]byte),
	}
}

func (m *MockS3Client) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	// Simulate storing the object
	if input.Body == nil {
		return nil, fmt.Errorf("no body provided")
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(input.Body)
	m.StoredObjects[*input.Key] = buf.Bytes()

	// Return a successful response
	return &s3.PutObjectOutput{}, nil
}

func (m *MockS3Client) DeleteObject(input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	// Simulate deleting the object
	if _, exists := m.StoredObjects[*input.Key]; !exists {
		return nil, fmt.Errorf("object not found")
	}

	delete(m.StoredObjects, *input.Key)

	// Return a successful response
	return &s3.DeleteObjectOutput{}, nil
}
