package store

import (
	"errors"
	"os"
	"path/filepath"
)

// GetExecPath gets the full path of the executable on the filesystem. This can be unreliable as
// it expects argv[0] to be the name of the executable which is not always the case. It also
// relies on the current working directory being the location of the executable.
func GetExecPath() (filename string, err error) {
	if len(os.Args) == 0 {
		return "", errors.New("useless: There are no arguments passed. Cannot get executable name.")
	}

	wd, err := os.Getwd()
	if err != nil {
		return
	}
	return filepath.Join(wd, os.Args[0]), nil
}
