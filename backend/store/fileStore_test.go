package store

import (
	"reflect"
	"testing"

	"github.com/paupenin/web2image/backend/config"
)

func TestNewFileStore(t *testing.T) {
	tests := []struct {
		name     string
		config   config.FileStoreConfig
		wantType reflect.Type
	}{
		{
			name:     "FileStoreFS",
			config:   &config.FileStoreFSConfig{},
			wantType: reflect.TypeOf(&FileStoreFS{}),
		},
		{
			name:     "FileStoreS3",
			config:   &config.FileStoreS3Config{},
			wantType: reflect.TypeOf(&FileStoreS3{}),
		},
		{
			name:     "FileStoreMemory (default)",
			config:   nil, // No specific config provided
			wantType: reflect.TypeOf(&FileStoreMemory{}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFileStore(tt.config)
			if reflect.TypeOf(got) != tt.wantType {
				t.Errorf("NewFileStore() = %v, want %v", reflect.TypeOf(got), tt.wantType)
			}
		})
	}
}
