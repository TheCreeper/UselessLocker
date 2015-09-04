// Package store attempts to open the currently running or specified executable
// as a zip archive.It also provides a means to easily find and access files
// within the archive.
package store

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

var ErrNotFound = errors.New("useless/store: File does not exist within the archive")

// GetExecPath gets the path of the executable on the filesystem. This can be
// unreliable as it expects argv[0] to be the name/path of the executable.
func GetExecPath() (string, error) {
	return exec.LookPath(os.Args[0])
}

type Store struct{ zr *zip.ReadCloser }

// Open will try to find the path of the executable and open a zip reader.
// This relies on argv[0] being the relative path to the executable and in most
// cases this is true.
func Open() (Store, error) {
	filename, err := GetExecPath()
	if err != nil {
		return Store{}, err
	}
	return OpenFile(filename)
}

// OpenFile will try to read the specified file regardless if its a executable
// or zip file.
func OpenFile(filename string) (Store, error) {
	file, err := zip.OpenReader(filename)
	if err != nil {
		return Store{}, err
	}
	return Store{file}, nil
}

// Close will close the underlaying zip reader.
func (s Store) Close() error {
	return s.zr.Close()
}

// Load will attempt to look for a file within the store matching the specified
// filepath relative to the store. If found the file contents are copied to
// memory.
func (s Store) Load(filename string) (fileBytes []byte, err error) {
	for _, file := range s.zr.File {
		if !file.FileInfo().IsDir() || file.Name != filename {
			continue
		}

		rc, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer rc.Close()

		fileBytes = make([]byte, file.FileInfo().Size())
		if _, err = rc.Read(fileBytes); err != nil {
			return nil, err
		}
	}
	return nil, ErrNotFound
}

// Unpack will copy the contents of the store to the specified location on the
// filesystem. No error wil be returned if there is nothing to be extracted
// from the store; it always assumed that there is something in the store to
// extracted.
func (s Store) UnPack(dst string) (err error) {
	for _, file := range s.zr.File {
		if file.FileInfo().IsDir() {
			if err = os.Mkdir(file.Name, 0755); err != nil {
				return
			}
			continue
		}

		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		nfile, err := os.Create(filepath.Join(dst, file.Name))
		if err != nil {
			return err
		}
		defer nfile.Close()

		if _, err = io.Copy(nfile, rc); err != nil {
			return err
		}
	}
	return
}
