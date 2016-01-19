package useless

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
)

// GetFileList compiles a list of files that match any of the extensions in the
// FileExtensions map and are less than or equal to the specified file size.
func GetFileList(dirname string, size int64) (files []string, err error) {
	list, err := ioutil.ReadDir(dirname)
	if err != nil {
		return
	}
	for _, file := range list {
		path := filepath.Join(dirname, file.Name())

		// We don't want to follow symbolic links as it may cause
		// some issues especially on windows.
		if file.IsDir() && (file.Mode()&os.ModeSymlink == 0) {
			fl, err := GetFileList(path, size)
			if err != nil {
				return nil, err
			}
			files = append(files, fl...)
			continue
		}

		if file.Mode().IsRegular() &&
			(file.Size() <= size) &&
			FileExtensions[filepath.Ext(file.Name())] {
			files = append(files, path)
		}
	}
	return
}

// WriteFileList writes a list of files to the specified directory.
func WriteFileList(dirname string, files []string) (err error) {
	// Add a newline character after every filepath
	var data []byte
	for _, v := range files {
		data = append(data, []byte(v)...)
		data = append(data, []byte(string('\n'))...)
	}

	path := filepath.Join(dirname, PathFileList)
	if err = ioutil.WriteFile(path, data, 0644); err != nil {
		return
	}
	return
}

// ReadFileList reads a list of files contained in a file within the users
// home directory.
func ReadFileList(dirname string) (files []string, err error) {
	b, err := ioutil.ReadFile(filepath.Join(dirname, PathFileList))
	if err != nil {
		return
	}

	buf := bytes.NewBuffer(b)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			return nil, err
		}
		files = append(files, line)
	}
	return
}
