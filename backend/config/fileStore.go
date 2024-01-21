package config

const (
	StorageTypeLocal = "local"
	StorageTypeS3    = "s3"
)

type FileStoreConfig interface {
	// The type of storage system (local, s3)
	Type() string
	// Whether the storage system should serve static files
	ShouldServeStatic() bool
}

// NewStorageConfig creates a new storage config
func NewFileStoreConfig() FileStoreConfig {
	storageType := Env("STORAGE_TYPE", StorageTypeLocal)

	if storageType == StorageTypeS3 {
		return &FileStoreS3Config{
			BucketName: Env("STORAGE_S3_BUCKET_NAME", ""),
			Region:     Env("STORAGE_S3_REGION", ""),
			AccessKey:  Env("STORAGE_S3_ACCESS_KEY", ""),
			SecretKey:  Env("STORAGE_S3_SECRET_KEY", ""),
		}
	}

	return &FileStoreFSConfig{
		Directory: Env("STORAGE_FS_DIRECTORY", "../store/images"),
		PublicURL: Env("STORAGE_FS_PUBLIC_URL", "http://localhost:8080/images"),
	}
}

// FileStoreFSConfig struct for FileSystem configuration
type FileStoreFSConfig struct {
	Directory string
	PublicURL string
}

func (c *FileStoreFSConfig) Type() string {
	return StorageTypeLocal
}

func (c *FileStoreFSConfig) ShouldServeStatic() bool {
	return true
}

// S3Config struct for S3 configuration
type FileStoreS3Config struct {
	BucketName string
	Region     string
	AccessKey  string
	SecretKey  string
}

func (c *FileStoreS3Config) Type() string {
	return StorageTypeS3
}

func (c *FileStoreS3Config) ShouldServeStatic() bool {
	return false
}
