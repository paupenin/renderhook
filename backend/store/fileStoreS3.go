package store

import (
	"bytes"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/paupenin/web2image/backend/config"
)

type S3API interface {
	PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
	DeleteObject(input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error)
}

type FileStoreS3 struct {
	config *config.FileStoreS3Config
	client S3API
}

func NewFileStoreS3(cfg *config.FileStoreS3Config) *FileStoreS3 {
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(cfg.Region),
		Endpoint:         aws.String(cfg.Endpoint),
		Credentials:      credentials.NewStaticCredentials(cfg.AccessKey, cfg.SecretKey, ""),
		S3ForcePathStyle: aws.Bool(cfg.S3ForcePathStyle),
		DisableSSL:       aws.Bool(!cfg.SSL),
	})
	if err != nil {
		log.Fatal("error creating AWS session: %w", err)
		return nil
	}

	return &FileStoreS3{
		config: cfg,
		client: s3.New(sess),
	}
}

func (s *FileStoreS3) ShouldServeStatic() bool {
	return false
}

func (s *FileStoreS3) GetStaticPath() string {
	return "" // Not applicable for S3
}

func (s *FileStoreS3) StoreFile(filename string, file []byte) error {
	_, err := s.client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(s.config.BucketName),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(file),
		ContentType: aws.String("application/octet-stream"), // Update as needed
	})
	return err
}

func (s *FileStoreS3) GetFileURL(filename string) string {
	return fmt.Sprintf("%s/%s", s.config.Endpoint, filename)
}

func (s *FileStoreS3) DeleteFile(filename string) error {
	_, err := s.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.config.BucketName),
		Key:    aws.String(filename),
	})
	return err
}
