// Package store attempts to open the currently running or specified executable as a zip file. It also provides a means to easily find and access files within the archive.
package store

import (
	"archive/zip"
	"errors"
	"os"
	"os/exec"
)

// Errors
var (
	ErrNoFile = errors.New("useless/store: File does not exist within archive")
)

// GetExecPath gets the path of the executable on the filesystem. This can be unreliable as
// it expects argv[0] to be the name/path of the executable.
func GetExecPath() (string, error) {
	return exec.LookPath(os.Args[0])
}

type Store struct{ zr *zip.ReadCloser }

// Open will try to find the path of the executable and open a zip reader. This relies on
// argv[0] being the relative path to the executable and in most cases this is true.
func Open() (Store, error) {
	filename, err := GetExecPath()
	if err != nil {
		return Store{}, err
	}
	return OpenFile(filename)
}

// OpenFile will try to read the specified file regardless if its a executable or zip file.
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

// Load will attempt to look for a file within the store matching the specified filename. A
// filepath must be relative to the file store.
func (s Store) Load(filename string) ([]byte, error) {
	for _, file := range s.zr.File {
		if file.Name != filename {
			continue
		}

		rc, err := file.Open()
		if err != nil {
			return nil, err
		}

		fileBytes := make([]byte, file.FileInfo().Size())
		if _, err := rc.Read(fileBytes); err != nil {
			return nil, err
		}
		return fileBytes, nil
	}
	return nil, ErrNoFile
}

// ExtractTo will look for the a filename/dir and extract the contents to the specified directory.
func (s Store) ExtractTo(filename string, to string) (err error) {
	return
}
