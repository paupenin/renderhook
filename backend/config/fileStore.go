package config

const (
	StorageTypeLocal = "local"
	StorageTypeS3    = "s3"
)

type FileStoreConfig interface{}

// NewStorageConfig creates a new storage config
func NewFileStoreConfig() FileStoreConfig {
	storageType := Env("STORAGE_TYPE", StorageTypeLocal)

	if storageType == StorageTypeS3 {
		return &FileStoreS3Config{
			BucketName:       Env("STORAGE_S3_BUCKET_NAME", ""),
			Region:           Env("STORAGE_S3_REGION", ""),
			AccessKey:        Env("STORAGE_S3_ACCESS_KEY", ""),
			SecretKey:        Env("STORAGE_S3_SECRET_KEY", ""),
			Endpoint:         Env("STORAGE_S3_ENDPOINT", ""),
			SSL:              EnvBool("STORAGE_S3_SSL", true),
			S3ForcePathStyle: EnvBool("STORAGE_S3_FORCE_PATH_STYLE", false),
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

// S3Config struct for S3 configuration
type FileStoreS3Config struct {
	BucketName       string
	Region           string
	AccessKey        string
	SecretKey        string
	Endpoint         string // Endpoint URL of the S3-compatible service
	SSL              bool   // Whether to use SSL (HTTPS)
	S3ForcePathStyle bool   // Use path style for bucket access
}
