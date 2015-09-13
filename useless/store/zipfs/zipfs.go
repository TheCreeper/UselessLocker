package zipfs

import (
	"archive/zip"
	"io"
	"os"
)

// File represents a file in the filesystem.
type File struct {
	io.Reader
	io.Closer
}

func (f File) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f File) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func (f File) Stat() (os.FileInfo, error) {
	return nil, nil
}

// FileSystem is the filesystem interface.
type FileSystem struct{ *zip.ReadCloser }

// Open opens a file, returning it or an error, if any happens.
func (fs FileSystem) Open(name string) (File, error) {
	for _, file := range fs.File {
		if file.FileInfo().IsDir() || file.Name != name {
			continue
		}

		rc, err := file.Open()
		if err != nil {
			return File{}, err
		}
		return File{rc, rc}, nil
	}
	return File{}, nil
}
