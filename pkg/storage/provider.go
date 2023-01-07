package storage

import (
	"github.com/casdoor/oss"
)

// Config - storage config
type Config struct {
	ID       string
	Secret   string
	Region   string
	Bucket   string
	Endpoint string
}

// GetProvider - get new storage provider
func GetProvider(provider string, c *Config) oss.StorageInterface {
	switch provider {
	case "local":
		return NewFileSystem(c.Bucket)
	case "s3":
		return NewS3(c)
	case "qiniu":
		return NewQiniu(c)
	case "oss":
		return NewAliyun(c)
	case "cos":
		return NewTencentCloud(c)
	case "azure":
		return NewAzure(c)
	}
	return nil
}
