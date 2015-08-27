//go:generate go-bindata -pkg=store -o=./bin.go -prefix=../../ ../../assets
package store

import (
	"archive/zip"
	"debug/elf"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Store struct{ TmpDir string }

// Open will attempt to find and extract a gzipped tar archive inside or appended to an executable.
func Open() (s Store, err error) {
	// Get the full path of the running exectable. This can be unreliable if argv[0] is
	// different than what is expected.
	filename, err := GetExecPath()
	if err != nil {
		return
	}

	// Open the executable with the read only flag
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	// Get the file size of the executable
	finfo, err := file.Stat()
	if err != nil {
		return
	}

	// Create a new temp directory to untar the files into
	tmpdir, err := ioutil.TempDir(os.TempDir(), "store")
	if err != nil {
		return
	}
	s.TmpDir = tmpdir

	// Look for a gzip file in the executable and return the gzip reader
	zr, err := elfZipReader(file, finfo.Size())
	if err != nil {
		return
	}
	for _, f := range zr.File {
		// Open a reader to the file within the archive
		rc, err := f.Open()
		if err != nil {
			return Store{}, err
		}
		defer rc.Close()

		// Create a file in the temp folder with the same name as in the tar archive
		nfile, err := os.Create(filepath.Join(s.TmpDir, f.Name))
		if err != nil {
			return Store{}, err
		}

		// Copy the bytes from the tar archive to the newly created file
		_, err = io.Copy(nfile, rc)
		if err != nil {
			return Store{}, err
		}
	}
	return
}

// Load will attempt to look for a file within the store matching the specified filename. A
// filepath must be relative to the file store.
func (s Store) Load(filename string) (b []byte, err error) {
	return ioutil.ReadFile(filepath.Join(s.TmpDir, filename))
}

func elfZipReader(r io.ReaderAt, size int64) (*zip.Reader, error) {
	file, err := elf.NewFile(r)
	if err != nil {
		return nil, err
	}

	var max int64
	for _, sect := range file.Sections {
		if sect.Type == elf.SHT_NOBITS {
			continue
		}

		zr, err := zip.NewReader(sect, int64(sect.Size))
		if err == nil {
			// There is a zip file here.
			return zr, nil
		}

		// Move end of file pointer
		end := int64(sect.Offset + sect.Size)
		if end > max {
			max = end
		}
	}

	// If zip archive not within binary. Check to see if its appended to the end.
	section := io.NewSectionReader(r, max, size-max)
	return zip.NewReader(section, section.Size())
}
