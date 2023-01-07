package storage

import (
	"github.com/casdoor/oss"
	"github.com/casdoor/oss/s3"
)

// NewS3 - new aws s3
func NewS3(c *Config) oss.StorageInterface {
	return s3.New(&s3.Config{
		AccessID:  c.ID,
		AccessKey: c.Secret,
		Region:    c.Region,
		Bucket:    c.Bucket,
		Endpoint:  c.Endpoint,
	})
}
