package storage

import (
	"github.com/casdoor/oss"
	"github.com/casdoor/oss/qiniu"
)

// NewQiniu - new qiniu storage
func NewQiniu(c *Config) oss.StorageInterface {
	return qiniu.New(&qiniu.Config{
		AccessID:  c.ID,
		AccessKey: c.Secret,
		Region:    c.Region,
		Bucket:    c.Bucket,
		Endpoint:  c.Endpoint,
	})
}
