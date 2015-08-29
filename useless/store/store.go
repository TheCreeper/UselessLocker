// Package store attempts to open the currently running or specified executable as a zip file. It also provides a means to easily find and access files within the archive.
package store

import (
	"archive/zip"
	"errors"
	"io"
)

// Errors
var (
	ErrNoFile = errors.New("useless/store: File does not exist within archive")
)

type Store struct{ zr *zip.ReadCloser }

// Open will attempt to find the path of the executable and open a zip reader. This relies on
// argv[0] being the relative path to the executable and in most cases this is true.
func Open() (Store, error) {
	// Get the full path of the running exectable. This can be unreliable if argv[0] is
	// different than what is expected.
	filename, err := GetExecPath()
	if err != nil {
		return Store{}, err
	}
	return OpenFile(filename)
}

// OpenFile will attempt to read the specified file regardless if its a executable or zip file.
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

		b := make([]byte, file.FileInfo().Size())
		if _, err := rc.Read(b); err != nil {
			return nil, err
		}
		return b, nil
	}
	return nil, ErrNoFile
}

// LoadReader behaves the same as Load expect that it returns a io.ReadCloser for the
// specified file.
func (s Store) LoadReader(filename string) (rc io.ReadCloser, err error) {
	for _, file := range s.zr.File {
		if file.Name != filename {
			continue
		}

		rc, err = file.Open()
		if err != nil {
			return
		}
		return
	}
	return nil, ErrNoFile
}

/*func elfGetSize(r io.ReaderAt) (size int64, err error) {
	file, err := elf.NewFile(r)
	if err != nil {
		return
	}
	for _, sect := range file.Sections {
		if sect.Type == elf.SHT_NOBITS {
			continue
		}

		// Move end of file pointer
		end := int64(sect.Offset + sect.Size)
		if end > size {
			size = end
		}
	}
	return
}*/
