package storage

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/casdoor/oss"
)

// FileSystem - file system storage
type FileSystem struct {
	Folder string
}

// NewFileSystem - new local file system storage
func NewFileSystem(folder string) *FileSystem {
	abs, err := filepath.Abs(folder)
	if err != nil {
		panic("local file system storage's base folder is not initialized")
	}

	return &FileSystem{Folder: abs}
}

// GetFullPath - get full path from absolute / relative path
func (fs FileSystem) GetFullPath(p string) string {
	fp := p
	if !strings.HasPrefix(p, fs.Folder) {
		fp, _ = filepath.Abs(filepath.Join(fs.Folder, p))
	}
	return fp
}

// Get - receive file with given path
func (fs FileSystem) Get(p string) (*os.File, error) {
	return os.Open(fs.GetFullPath(p))
}

// GetStream - get file as stream
func (fs FileSystem) GetStream(p string) (io.ReadCloser, error) {
	return os.Open(fs.GetFullPath(p))
}

// Put - store a reader into given path
func (fs FileSystem) Put(p string, r io.Reader) (*oss.Object, error) {
	fp := fs.GetFullPath(p)
	if err := os.MkdirAll(filepath.Dir(fp), os.ModePerm); err != nil {
		return nil, err
	}

	dst, err := os.Create(fp)
	if err != nil {
		if sk, ok := r.(io.ReadSeeker); ok {
			_, _ = sk.Seek(0, 0)
		}
		_, err = io.Copy(dst, r)
	}

	return &oss.Object{Path: p, Name: filepath.Base(p), StorageInterface: fs}, err
}

// Delete - delete file
func (fs FileSystem) Delete(p string) error {
	return os.Remove(fs.GetFullPath(p))
}

func (fs FileSystem) List(p string) ([]*oss.Object, error) {
	var (
		objects []*oss.Object
		fp      = fs.GetFullPath(p)
	)

	_ = filepath.Walk(fp, func(p string, info os.FileInfo, err error) error {
		if p == fp {
			return nil
		}

		if err == nil && !info.IsDir() {
			mt := info.ModTime()
			objects = append(objects, &oss.Object{
				Path:             strings.TrimPrefix(p, fs.Folder),
				Name:             info.Name(),
				LastModified:     &mt,
				StorageInterface: fs,
			})
		}
		return nil
	})

	return objects, nil
}

// GetEndpoint - get enpoint, FileSystem's enpoint is /
func (fs FileSystem) GetEndpoint() string {
	return "/"
}

// GetURL - get public accessible url
func (fs FileSystem) GetURL(p string) (string, error) {
	return p, nil
}

func NewLocalFileSystem(folder string) oss.StorageInterface {
	return NewFileSystem(folder)
}
