package store

import (
	"archive/zip"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/httpfs"
	"golang.org/x/tools/godoc/vfs/zipfs"
)

type StoreFS struct{ vfs.FileSystem }

// Open will try to find the path of the executable and open a zip reader.
// This relies on argv[0] being the relative path to the executable and in most
// cases this is true.
func Open() (StoreFS, error) {
	// This can be unreliable as it expects argv[0] to be the name/path of
	// the executable.
	filename, err := exec.LookPath(os.Args[0])
	if err != nil {
		return StoreFS{}, err
	}
	return OpenFile(filename)
}

func OpenFile(filename string) (StoreFS, error) {
	rc, err := zip.OpenReader(filename)
	if err != nil {
		return StoreFS{}, err
	}
	return StoreFS{zipfs.New(rc, filename)}, nil
}

func (fs StoreFS) ReadFile(filename string) (b []byte, err error) {
	file, err := fs.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

func (fs StoreFS) Dir(filename string) http.FileSystem {
	return httpfs.New(fs)
}
