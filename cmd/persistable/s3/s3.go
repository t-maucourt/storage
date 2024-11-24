package s3

import (
	"log"
)

type s3Storage struct {
	bucket string
}

func NewS3Storage(bucket string) *s3Storage {
	log.Printf("New S3 Storage targeting bucket %s\n", bucket)
	return &s3Storage{bucket}
}

func (s *s3Storage) Save(b []byte, args ...any) error {
	log.Printf("Saving to S3 (%s : %s): %s\n", s.bucket, args, b)
	return nil
}

func (s *s3Storage) Load(args ...any) ([]byte, error) {
	log.Printf("Loading from S3 (%s) - %s\n", s.bucket, args)
	return []byte{}, nil
}
