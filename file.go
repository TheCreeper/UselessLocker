package useless

import (
	"bufio"
	"os"
	"path/filepath"
)

// GetFileList compiles a list of files that match any of the extensions in the
// FileExtensions map and are less than or equal to the specified file size.
func GetFileList(dirname string, size int64) (files []string, err error) {
	walkFn := func(path string, info os.FileInfo, err error) error {
		// Can't walk here.
		if err != nil {
			// Continue walking elsewhere.
			return nil
		}

		// Check if this is a regular file otherwise skip it.
		if !info.Mode().IsRegular() {
			return nil
		}

		// Check if this has a size equal or less than to the specified
		// file size.
		if !(info.Size() <= size) {
			return nil
		}

		// Check if this has a file extension that is in the
		// FileExtensions map.
		if !FileExtensions[filepath.Ext(info.Name())] {
			return nil
		}

		files = append(files, path)
		return nil
	}

	err = filepath.Walk(dirname, walkFn)
	return
}

// WriteFileList writes a list of files to the specified directory.
func WriteFileList(filename string, files []string) (err error) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	for _, filepath := range files {
		_, err = f.WriteString(filepath + "\n")
		if err != nil {
			return
		}
	}
	return
}

// ReadFileList reads a list of files contained in a file within the users
// home directory.
func ReadFileList(filename string) (files []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Append each line of the file list to the files string
		// slice.
		files = append(files, scanner.Text())
	}

	err = scanner.Err()
	return
}
