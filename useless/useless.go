package useless

import (
	"io/ioutil"
	"os/user"
	"path/filepath"
)

func Start() (err error) {
	user, err := user.Current()
	if err != nil {
		return
	}

	// Get a list of files in the users home directory that are less than 10MB
	files, err := GetFileList(user.HomeDir, 10485760)
	if err != nil {
		return
	}
	return
}

// GetFileList compiles a list of files that are less than or equal to the specified file size.
// The file list is returned as a string slice with full location path.
func GetFileList(dirname string, size int64) (files []string, err error) {
	list, err := ioutil.ReadDir(dirname)
	if err != nil {
		return
	}
	for _, v := range list {
		if v.IsDir() {
			fl, err := GetFileList(filepath.Join(dirname, v.Name()), size)
			if err != nil {
				return nil, err
			}
			files = append(files, fl...)
			continue
		}

		if v.Mode().IsRegular() &&
			(v.Mode().Perm() <= 0777) &&
			(v.Size() <= size) {
			files = append(files, filepath.Join(dirname, v.Name()))
		}
	}
	return
}
